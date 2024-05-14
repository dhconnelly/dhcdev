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
	"strings"

	"github.com/yuin/goldmark"
)

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

type page struct {
	Title   string
	Content template.HTML
}

func renderPage(w io.Writer, r io.Reader, tmpl *template.Template) error {
	br := bufio.NewReader(r)
	title, err := br.ReadBytes('\n')
	if err != nil {
		return err
	}
	content, err := io.ReadAll(br)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err = goldmark.Convert(content, &buf); err != nil {
		return err
	}
	return tmpl.Execute(w, page{
		Title:   string(title),
		Content: template.HTML(buf.String()),
	})
}

func renderFile(outPath, srcPath string, tmpl *template.Template) error {
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
	return renderPage(to, from, tmpl)
}

func shouldRender(srcPath string) bool {
	return filepath.Ext(srcPath) == ".md"
}

func changeExtension(path string, ext string) string {
	filename := strings.TrimSuffix(path, filepath.Ext(path))
	return fmt.Sprintf("%s.%s", filename, ext)
}

func buildPage(path, srcDir, outDir string, tmpl *template.Template) error {
	if shouldRender(path) {
		srcPath := filepath.Join(srcDir, path)
		outPath := filepath.Join(outDir, changeExtension(path, "html"))
		return renderFile(outPath, srcPath, tmpl)
	} else {
		srcPath := filepath.Join(srcDir, path)
		outPath := filepath.Join(outDir, path)
		return copyFile(outPath, srcPath)
	}
}

func buildSiteWithTemplate(outDir, srcDir string, tmpl *template.Template) error {
	log.Printf("building from %s to %s", srcDir, outDir)
	return fs.WalkDir(
		os.DirFS(srcDir), ".",
		func(path string, d fs.DirEntry, err error) error {
			outPath := filepath.Join(outDir, path)
			if d.IsDir() {
				return os.MkdirAll(outPath, 0755)
			} else {
				return buildPage(path, srcDir, outDir, tmpl)
			}
		})
}

func buildSite(outDir, srcDir, tmplPath string) error {
	log.Println("building site...")
	log.Printf("loading template %s", tmplPath)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	return buildSiteWithTemplate(outDir, srcDir, tmpl)
}
