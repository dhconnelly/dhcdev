package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

type page struct {
	Title   string
	Content template.HTML
}

var titlePat = regexp.MustCompile(`^=== ([^=]+) ===$`)

func readTitle(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading title: %w", err)
	}
	ms := titlePat.FindStringSubmatch(line[:len(line)-1])
	if len(ms) != 2 {
		return "", fmt.Errorf("title is missing")
	}
	return ms[1], nil
}

type builder struct {
	srcDir string
	outDir string
	md     goldmark.Markdown
	tmpl   *template.Template
}

func newBuilder(srcDir, outDir string, tmpl *template.Template) *builder {
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
					chromahtml.TabWidth(4),
				),
			),
		),
	)
	return &builder{srcDir: srcDir, outDir: outDir, md: md, tmpl: tmpl}
}

func (b *builder) renderPage(w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)
	title, err := readTitle(br)
	if err != nil {
		return err
	}
	content, err := io.ReadAll(br)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err = b.md.Convert(content, &buf); err != nil {
		return err
	}
	return b.tmpl.Execute(w, page{
		Title:   string(title),
		Content: template.HTML(buf.String()),
	})
}

func (b *builder) renderFile(outPath, srcPath string) error {
	log.Printf("rendering %s to %s", srcPath, outPath)
	from, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer from.Close()
	to, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer to.Close()
	return b.renderPage(to, from)
}

func copyFile(outPath, srcPath string) error {
	log.Printf("copying %s to %s", srcPath, outPath)
	from, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer from.Close()
	to, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer to.Close()
	_, err = io.Copy(to, from)
	return err
}

func shouldRender(srcPath string) bool {
	return filepath.Ext(srcPath) == ".md"
}

func changeExtension(path string, ext string) string {
	filename := strings.TrimSuffix(path, filepath.Ext(path))
	return fmt.Sprintf("%s.%s", filename, ext)
}

func (b *builder) buildPage(path string) error {
	if shouldRender(path) {
		srcPath := filepath.Join(b.srcDir, path)
		outPath := filepath.Join(b.outDir, changeExtension(path, "html"))
		return b.renderFile(outPath, srcPath)
	} else {
		srcPath := filepath.Join(b.srcDir, path)
		outPath := filepath.Join(b.outDir, path)
		return copyFile(outPath, srcPath)
	}
}

func BuildSite(outDir, srcDir, tmplPath string) error {
	log.Println("building site...")
	log.Printf("loading template %s", tmplPath)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	log.Printf("building from %s to %s", srcDir, outDir)
	b := newBuilder(srcDir, outDir, tmpl)
	return fs.WalkDir(
		os.DirFS(srcDir), ".",
		func(path string, d fs.DirEntry, err error) error {
			outPath := filepath.Join(outDir, path)
			if d.IsDir() {
				return os.MkdirAll(outPath, 0755)
			} else {
				return b.buildPage(path)
			}
		})
}
