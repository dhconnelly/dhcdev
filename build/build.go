package main

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

func buildFile(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
}

func visitor(dstDir, srcDir string) fs.WalkDirFunc {
	return func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println("error:", err)
			return nil // keep going
		}
		if d.IsDir() {
			return nil // keep going
		}

		// prep the output
		relPath, err := filepath.Rel(srcDir, filePath)
		if err != nil {
			log.Println("failed to get relative destination path:", err)
			return nil // keep going
		}
		dstPath := path.Join(dstDir, relPath)
		if err := os.MkdirAll(path.Dir(dstPath), 0755); err != nil {
			log.Println("error:", err)
			return nil // keep going
		}

		// open the files
		src, err := os.Open(filePath)
		if err != nil {
			log.Println("error opening src:", err)
			return nil // keep going
		}
		defer src.Close()

		// don't copy the src prefix
		dst, err := os.Create(dstPath)
		if err != nil {
			log.Println("error opening dst:", err)
			return nil // keep going
		}
		defer dst.Close()

		// build
		if err := buildFile(dst, src); err != nil {
			log.Printf("error building file %s: %s", filePath, err)
		}
		return nil // keep going
	}
}

func buildTree(dst, src string) error {
	log.Printf("building files from %s into %s", dst, src)
	fs.WalkDir(os.DirFS("."), src, visitor(dst, src))
	return nil
}
