use axum::extract::State;
use axum::http::{Request, StatusCode};
use axum::middleware::{self, Next};
use axum::response::{IntoResponse, Response};
use axum::routing::get;
use axum::Router;
use serde::Serialize;
use std::path::Path;
use std::sync::{Arc, Mutex};
use std::time::Instant;
use tower_http::services::ServeDir;

#[derive(Default, Serialize)]
struct Metrics {
    requests_total: u64,
    requests_success: u64,
    requests_client_error: u64,
    requests_server_error: u64,
    request_duration_ms_sum: u64,
    request_duration_ms_count: u64,
}

#[derive(Clone)]
pub struct AppState {
    metrics: Arc<Mutex<Metrics>>,
}

async fn healthz() -> StatusCode {
    StatusCode::OK
}

async fn varz(State(state): State<AppState>) -> impl IntoResponse {
    let metrics = state.metrics.lock().unwrap();
    (
        StatusCode::OK,
        [("content-type", "application/json")],
        serde_json::to_string_pretty(&*metrics).unwrap(),
    )
}

async fn logging_middleware(
    State(state): State<AppState>,
    request: Request<axum::body::Body>,
    next: Next,
) -> Response {
    let method = request.method().clone();
    let uri = request.uri().clone();
    let start = Instant::now();

    let response = next.run(request).await;

    let status = response.status();
    let duration_ms = start.elapsed().as_millis() as u64;

    println!("\"{method} {uri} HTTP/1.1\" {status} {duration_ms}ms");

    let mut metrics = state.metrics.lock().unwrap();
    metrics.requests_total += 1;
    if status.is_success() {
        metrics.requests_success += 1;
    } else if status.is_client_error() {
        metrics.requests_client_error += 1;
    } else if status.is_server_error() {
        metrics.requests_server_error += 1;
    }
    metrics.request_duration_ms_sum += duration_ms;
    metrics.request_duration_ms_count += 1;

    response
}

pub fn app(serve_dir: &Path) -> Router {
    let state = AppState {
        metrics: Arc::new(Mutex::new(Metrics::default())),
    };

    Router::new()
        .route("/healthz", get(healthz))
        .route("/varz", get(varz))
        .fallback_service(ServeDir::new(serve_dir))
        .layer(middleware::from_fn_with_state(
            state.clone(),
            logging_middleware,
        ))
        .with_state(state)
}

pub async fn serve(serve_dir: &Path, port: u16) {
    let app = app(serve_dir);
    println!("serving {serve_dir:?} at http://localhost:{port}");
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    axum::serve(listener, app).await.unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;
    use axum::body::Body;
    use axum::http::Request;
    use tempfile::TempDir;
    use tower::ServiceExt;

    async fn request(app: Router, uri: &str) -> Response {
        app.oneshot(
            Request::builder()
                .uri(uri)
                .body(Body::empty())
                .unwrap(),
        )
        .await
        .unwrap()
    }

    #[tokio::test]
    async fn test_healthz() {
        let tmp = TempDir::new().unwrap();
        let response = request(app(tmp.path()), "/healthz").await;
        assert_eq!(response.status(), StatusCode::OK);
    }

    #[tokio::test]
    async fn test_not_found() {
        let tmp = TempDir::new().unwrap();
        let response = request(app(tmp.path()), "/nonexistent").await;
        assert_eq!(response.status(), StatusCode::NOT_FOUND);
    }

    #[tokio::test]
    async fn test_static_file() {
        let tmp = TempDir::new().unwrap();
        std::fs::write(tmp.path().join("hello.txt"), "hello world").unwrap();

        let response = request(app(tmp.path()), "/hello.txt").await;
        assert_eq!(response.status(), StatusCode::OK);
        let body = axum::body::to_bytes(response.into_body(), usize::MAX)
            .await
            .unwrap();
        assert_eq!(&body[..], b"hello world");
    }

    #[tokio::test]
    async fn test_varz_json() {
        let tmp = TempDir::new().unwrap();
        let response = request(app(tmp.path()), "/varz").await;
        assert_eq!(response.status(), StatusCode::OK);
        let body = axum::body::to_bytes(response.into_body(), usize::MAX)
            .await
            .unwrap();
        let v: serde_json::Value = serde_json::from_slice(&body).unwrap();
        assert_eq!(v["requests_total"], 0);
    }
}
