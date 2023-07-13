package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/russross/blackfriday"
)

var titlePat = regexp.MustCompile(`^=== ([^=]+) ===$`)

func readTitle(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading title: %w", err)
	}
	ms := titlePat.FindStringSubmatch(line[:len(line)-1])
	if len(ms) != 2 {
		return "", errors.New("title is missing")
	}
	return ms[1], nil
}

func buildMarkdown(dst io.Writer, src io.Reader) error {
	r := bufio.NewReader(src)
	w := bufio.NewWriter(dst)
	// TODO: build a proper page

	// process the title
	title, err := readTitle(r)
	if err != nil {
		return fmt.Errorf("error building markdown: %w", err)
	}
	if _, err = w.WriteString(fmt.Sprintf("<head><title>%s</title></head>", title)); err != nil {
		return fmt.Errorf("error writing title: %w", err)
	}

	// process the content
	input, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("error reading source: %w", err)
	}
	output := blackfriday.MarkdownCommon(input)
	if _, err = w.Write(output); err != nil {
		return fmt.Errorf("error writing content: %w", err)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("error writing title: %w", err)
	}
	return nil
}

func BuildFile(path string, dst io.Writer, src io.Reader) error {
	if strings.HasSuffix(path, "md") {
		return buildMarkdown(dst, src)
	}
	_, err := io.Copy(dst, src)
	return err
}

func makeDstPath(srcDir, dstDir, filePath string) (string, error) {
	relPath, err := filepath.Rel(srcDir, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative dest path: %w", err)
	}
	if strings.HasSuffix(relPath, "md") {
		relPath = strings.TrimSuffix(relPath, filepath.Ext(relPath))
		relPath = relPath + ".html"
	}
	dstPath := path.Join(dstDir, relPath)
	if err := os.MkdirAll(path.Dir(dstPath), 0755); err != nil {
		return "", fmt.Errorf("failed to make dest dirs: %w", err)
	}
	return dstPath, nil
}

func walk(dstDir, srcDir string) fs.WalkDirFunc {
	return func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			// io error unrelated to building
			return err
		}
		if d.IsDir() {
			return nil
		}

		dstPath, err := makeDstPath(srcDir, dstDir, filePath)
		if err != nil {
			return fmt.Errorf("error creating dest path: %w", err)
		}

		src, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("error opening src: %w", err)
		}
		defer src.Close()

		dst, err := os.Create(dstPath)
		if err != nil {
			return fmt.Errorf("error opening dst: %w", err)
		}
		defer dst.Close()

		if err := BuildFile(filePath, dst, src); err != nil {
			return fmt.Errorf("error building file %s: %s", filePath, err)
		}

		return nil
	}
}

func BuildTree(dst, src string) error {
	return filepath.WalkDir(src, walk(dst, src))
}
