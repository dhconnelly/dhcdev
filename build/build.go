package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

func BuildFile(dst io.Writer, src io.Reader) error {
	// TODO: handle headers
	// TODO: handle templates
	// TODO: build markdown files
	_, err := io.Copy(dst, src)
	return err
}

func makeDstPath(srcDir, dstDir, filePath string) (string, error) {
	relPath, err := filepath.Rel(srcDir, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative dest path: %w", err)
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
			// unknown i/o error: abort
			log.Println("error:", err)
			return err
		}

		if d.IsDir() {
			return nil // keep going
		}

		dstPath, err := makeDstPath(srcDir, dstDir, filePath)
		if err != nil {
			log.Println("error creating dest path:", err)
			return nil // keep going
		}

		src, err := os.Open(filePath)
		if err != nil {
			log.Println("error opening src:", err)
			return nil // keep going
		}
		defer src.Close()

		dst, err := os.Create(dstPath)
		if err != nil {
			log.Println("error opening dst:", err)
			return nil // keep going
		}
		defer dst.Close()

		if err := BuildFile(dst, src); err != nil {
			log.Printf("error building file %s: %s", filePath, err)
		}

		return nil // keep going
	}
}

func BuildTree(dst, src string) error {
	log.Printf("building files from %s into %s", src, dst)
	return filepath.WalkDir(src, walk(dst, src))
}
