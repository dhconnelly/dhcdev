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
    serve_until(serve_dir, port, shutdown_signal()).await;
}

async fn shutdown_signal() {
    use tokio::signal::unix::{signal, SignalKind};

    let mut sigint = signal(SignalKind::interrupt()).unwrap();
    let mut sigterm = signal(SignalKind::terminate()).unwrap();

    tokio::select! {
        _ = sigint.recv() => {}
        _ = sigterm.recv() => {}
    }
}

async fn serve_until(
    serve_dir: &Path,
    port: u16,
    shutdown: impl std::future::Future<Output = ()> + Send + 'static,
) {
    let app = app(serve_dir);
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    let addr = listener.local_addr().unwrap();
    println!("serving {serve_dir:?} at http://localhost:{}", addr.port());
    axum::serve(listener, app)
        .with_graceful_shutdown(shutdown)
        .await
        .unwrap();
    println!("server shut down");
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

    #[tokio::test]
    async fn test_graceful_shutdown() {
        use tokio::io::{AsyncReadExt, AsyncWriteExt};

        let tmp = TempDir::new().unwrap();
        std::fs::write(tmp.path().join("hello.txt"), "hi").unwrap();

        // Bind first so we know the port
        let listener = tokio::net::TcpListener::bind("127.0.0.1:0").await.unwrap();
        let port = listener.local_addr().unwrap().port();

        let (shutdown_tx, shutdown_rx) = tokio::sync::oneshot::channel::<()>();

        let path = tmp.path().to_path_buf();
        let server = tokio::spawn(async move {
            let app = app(&path);
            axum::serve(listener, app)
                .with_graceful_shutdown(async { shutdown_rx.await.ok(); })
                .await
                .unwrap();
        });

        // Verify server is responding
        let mut conn = tokio::net::TcpStream::connect(format!("127.0.0.1:{port}"))
            .await
            .unwrap();
        conn.write_all(b"GET /healthz HTTP/1.1\r\nHost: localhost\r\n\r\n")
            .await
            .unwrap();
        let mut buf = vec![0u8; 1024];
        let n = conn.read(&mut buf).await.unwrap();
        let response = String::from_utf8_lossy(&buf[..n]);
        assert!(response.contains("200 OK"));

        // Signal shutdown
        shutdown_tx.send(()).unwrap();

        // Server should exit cleanly
        tokio::time::timeout(std::time::Duration::from_secs(5), server)
            .await
            .expect("server did not shut down in time")
            .unwrap();
    }
}
