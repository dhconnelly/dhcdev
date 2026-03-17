mod render;
mod serve;

use std::env;
use std::path::PathBuf;

#[derive(Debug)]
struct Config {
    port: u16,
    source_dir: PathBuf,
    serve_dir: PathBuf,
    build: bool,
    serve: bool,
}

impl Config {
    fn from_env() -> Self {
        Config {
            port: env::var("PORT")
                .ok()
                .and_then(|v| v.parse().ok())
                .unwrap_or(9000),
            source_dir: env::var("SOURCE_DIR")
                .unwrap_or_else(|_| "./pages".into())
                .into(),
            serve_dir: env::var("SERVE_DIR")
                .unwrap_or_else(|_| "./out".into())
                .into(),
            build: env::var("BUILD")
                .map(|v| v.eq_ignore_ascii_case("true"))
                .unwrap_or(true),
            serve: env::var("SERVE")
                .map(|v| v.eq_ignore_ascii_case("true"))
                .unwrap_or(true),
        }
    }
}

async fn run(config: &Config) -> Result<(), render::RenderError> {
    println!("{config:?}");

    if config.build {
        render::build(&config.source_dir, &config.serve_dir)?;
    }

    if config.serve {
        serve::serve(&config.serve_dir, config.port).await;
    }

    Ok(())
}

#[tokio::main]
async fn main() {
    run(&Config::from_env()).await.unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;
    use axum::body::Body;
    use axum::http::{Request, StatusCode};
    use tempfile::TempDir;
    use tower::ServiceExt;

    fn test_config(source_dir: PathBuf, serve_dir: PathBuf) -> Config {
        Config {
            port: 0,
            source_dir,
            serve_dir,
            build: true,
            serve: false,
        }
    }

    #[tokio::test]
    async fn test_build_phase() {
        let src = TempDir::new().unwrap();
        let dst = TempDir::new().unwrap();

        std::fs::write(src.path().join("index.md"), "=== Home ===\n\nHello.").unwrap();
        std::fs::create_dir_all(src.path().join("css")).unwrap();
        std::fs::write(src.path().join("css/style.css"), "body {}").unwrap();

        run(&test_config(src.keep(), dst.path().into()))
            .await
            .unwrap();

        let html = std::fs::read_to_string(dst.path().join("index.html")).unwrap();
        assert!(html.contains("<title>Home</title>"));
        assert!(html.contains("Hello."));
        assert_eq!(
            std::fs::read_to_string(dst.path().join("css/style.css")).unwrap(),
            "body {}"
        );
    }

    #[tokio::test]
    async fn test_build_skipped_when_disabled() {
        let src = TempDir::new().unwrap();
        let dst = TempDir::new().unwrap();

        std::fs::write(src.path().join("index.md"), "=== Home ===\n\nHello.").unwrap();

        let config = Config {
            build: false,
            ..test_config(src.keep(), dst.path().into())
        };
        run(&config).await.unwrap();

        assert!(!dst.path().join("index.html").exists());
    }

    #[tokio::test]
    async fn test_build_then_serve() {
        let src = TempDir::new().unwrap();
        let dst = TempDir::new().unwrap();
        let dst_path = dst.path().to_path_buf();

        std::fs::write(src.path().join("index.md"), "=== Home ===\n\nHello.").unwrap();

        run(&test_config(src.keep(), dst_path.clone()))
            .await
            .unwrap();

        let resp = serve::app(&dst_path)
            .oneshot(Request::get("/healthz").body(Body::empty()).unwrap())
            .await
            .unwrap();
        assert_eq!(resp.status(), StatusCode::OK);

        let resp = serve::app(&dst_path)
            .oneshot(Request::get("/index.html").body(Body::empty()).unwrap())
            .await
            .unwrap();
        assert_eq!(resp.status(), StatusCode::OK);
        let body = axum::body::to_bytes(resp.into_body(), usize::MAX)
            .await
            .unwrap();
        let html = String::from_utf8(body.to_vec()).unwrap();
        assert!(html.contains("<title>Home</title>"));
        assert!(html.contains("Hello."));
    }
}
