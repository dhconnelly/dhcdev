use comrak::plugins::syntect::SyntectAdapterBuilder;
use comrak::{markdown_to_html_with_plugins, Options, Plugins};
use minijinja::{context, Environment};
use regex::Regex;
use std::fs;
use std::path::Path;
use std::sync::LazyLock;
use syntect::highlighting::ThemeSet;
use syntect::html::{css_for_theme_with_class_style, ClassStyle};

#[derive(Debug, thiserror::Error)]
pub enum RenderError {
    #[error("invalid page: {0}")]
    InvalidPage(String),
    #[error("{0}")]
    Io(#[from] std::io::Error),
    #[error("{0}")]
    Template(#[from] minijinja::Error),
}

const PAGE_TEMPLATE: &str = r#"<!DOCTYPE html>
<html>
    <head>
        <title>{{ title }}</title>
        <meta charset="utf-8" />
        <meta http-equiv="X-UA-Compatible" content="IE=edge" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta property="og:title" content="{{ title }}" />
        <meta name="author" content="Daniel Connelly" />
        <meta property="og:locale" content="en_US" />
        <meta property="og:type" content="website" />
        <meta name="twitter:card" content="summary" />
        <meta property="twitter:title" content="{{ title }}" />
        <link rel="stylesheet" type="text/css" href="/css/main.css">
        <link rel="stylesheet" type="text/css" href="/css/syntax.css">
    </head>
    <body>
        <main><div id="content">{{ content }}</div></main>
    </body>
</html>"#;

static TITLE_RE: LazyLock<Regex> =
    LazyLock::new(|| Regex::new(r"^=== ([^=]+) ===$").unwrap());

pub fn template_env() -> Environment<'static> {
    let mut env = Environment::new();
    env.add_template("page", PAGE_TEMPLATE).unwrap();
    env
}

pub fn generate_syntax_css() -> String {
    let theme_set = ThemeSet::load_defaults();
    let theme = &theme_set.themes["base16-eighties.dark"];
    css_for_theme_with_class_style(theme, ClassStyle::Spaced).unwrap()
}

fn render_markdown(source: &str) -> String {
    let adapter = SyntectAdapterBuilder::new().css().build();
    let mut options = Options::default();
    options.extension.header_ids = Some(String::new());
    let mut plugins = Plugins::default();
    plugins.render.codefence_syntax_highlighter = Some(&adapter);
    markdown_to_html_with_plugins(source, &options, &plugins)
}

pub fn render_page(
    env: &Environment,
    source: &str,
) -> Result<String, RenderError> {
    let first_line = source
        .lines()
        .next()
        .ok_or_else(|| RenderError::InvalidPage("empty file".into()))?;
    let title = TITLE_RE
        .captures(first_line)
        .and_then(|c| c.get(1))
        .map(|m| m.as_str())
        .ok_or_else(|| RenderError::InvalidPage("failed to parse title".into()))?;

    let rest = &source[first_line.len()..].trim_start_matches('\n');
    let content = render_markdown(rest);

    let tmpl = env.get_template("page").unwrap();
    Ok(tmpl.render(context! { title, content })?)
}

pub fn build(source_dir: &Path, dest_dir: &Path) -> Result<(), RenderError> {
    let env = template_env();
    println!("building {source_dir:?} to {dest_dir:?}");

    let css_dir = dest_dir.join("css");
    fs::create_dir_all(&css_dir)?;
    fs::write(css_dir.join("syntax.css"), generate_syntax_css())?;

    for entry in walk(source_dir)? {
        let rel = entry.strip_prefix(source_dir).unwrap();

        if entry.extension().and_then(|e| e.to_str()) == Some("md") {
            let mut dest = dest_dir.join(rel);
            dest.set_extension("html");
            println!("rendering {entry:?} to {dest:?}");
            if let Some(parent) = dest.parent() {
                fs::create_dir_all(parent)?;
            }
            let source = fs::read_to_string(&entry)?;
            let html = render_page(&env, &source)?;
            fs::write(&dest, html)?;
        } else {
            let dest = dest_dir.join(rel);
            println!("copying {entry:?} to {dest:?}");
            if let Some(parent) = dest.parent() {
                fs::create_dir_all(parent)?;
            }
            fs::copy(&entry, &dest)?;
        }
    }

    println!("built {source_dir:?} to {dest_dir:?}");
    Ok(())
}

fn walk(dir: &Path) -> std::io::Result<Vec<std::path::PathBuf>> {
    let mut files = Vec::new();
    walk_recursive(dir, &mut files)?;
    Ok(files)
}

fn walk_recursive(dir: &Path, files: &mut Vec<std::path::PathBuf>) -> std::io::Result<()> {
    for entry in fs::read_dir(dir)? {
        let entry = entry?;
        let path = entry.path();
        if path.is_dir() {
            walk_recursive(&path, files)?;
        } else {
            files.push(path);
        }
    }
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;
    use tempfile::TempDir;

    fn test_env() -> Environment<'static> {
        template_env()
    }

    fn test_render(source: &str) -> Result<String, RenderError> {
        let env = test_env();
        render_page(&env, source)
    }

    #[test]
    fn test_parse_title() {
        let html = test_render("=== My Title ===\n\nHello world").unwrap();
        assert!(html.contains("<title>My Title</title>"));
    }

    #[test]
    fn test_parse_title_fails_without_title() {
        assert!(test_render("no title here\n\nHello world").is_err());
    }

    #[test]
    fn test_markdown_rendering() {
        let html =
            test_render("=== Test ===\n\n# Heading\n\nA paragraph with **bold**.").unwrap();
        assert!(html.contains("<h1"));
        assert!(html.contains("Heading"));
        assert!(html.contains("<strong>bold</strong>"));
    }

    #[test]
    fn test_heading_anchors() {
        let html = test_render("=== Test ===\n\n## My Section").unwrap();
        assert!(html.contains(r#"id="my-section""#));
    }

    #[test]
    fn test_template_structure() {
        let html = test_render("=== Page Title ===\n\nContent here.").unwrap();
        assert!(html.contains("<!DOCTYPE html>"));
        assert!(html.contains(r#"<meta charset="utf-8" />"#));
        assert!(html.contains(r#"href="/css/main.css""#));
        assert!(html.contains(r#"<div id="content">"#));
        assert!(html.contains("Daniel Connelly"));
    }

    #[test]
    fn test_syntax_highlighting() {
        let env = test_env();
        let source = "=== Test ===\n\n```go\nfunc main() {}\n```";
        let html = render_page(&env, source).unwrap();
        assert!(html.contains("syntax-highlighting"));
        assert!(html.contains("class=\""));
    }

    #[test]
    fn test_syntax_css_generation() {
        let css = generate_syntax_css();
        assert!(css.contains(".code"));
    }

    #[test]
    fn test_build_pipeline() {
        let src = TempDir::new().unwrap();
        let dst = TempDir::new().unwrap();

        fs::write(
            src.path().join("index.md"),
            "=== Home ===\n\nWelcome home.",
        )
        .unwrap();

        fs::create_dir_all(src.path().join("css")).unwrap();
        fs::write(src.path().join("css/main.css"), "body { color: white; }").unwrap();

        build(src.path(), dst.path()).unwrap();

        let html = fs::read_to_string(dst.path().join("index.html")).unwrap();
        assert!(html.contains("<title>Home</title>"));
        assert!(html.contains("Welcome home."));

        let css = fs::read_to_string(dst.path().join("css/main.css")).unwrap();
        assert_eq!(css, "body { color: white; }");

        let syntax_css = fs::read_to_string(dst.path().join("css/syntax.css")).unwrap();
        assert!(syntax_css.contains(".code"));
    }
}
