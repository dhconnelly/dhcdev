package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var (
	srcDir   = flag.String("srcDir", "pages", "source directory to build")
	dstDir   = flag.String("dstDir", "target", "build output directory")
	postTmpl = flag.String("postTmpl", "templates/post-template.html", "path to the post template html")
	// TODO: inline the css
)

func main() {
	flag.Parse()
	log.Println("source directory:", *srcDir)
	log.Println("output directory:", *dstDir)
	log.Println("post template:", *postTmpl)
	tmplPath, err := filepath.Abs(*postTmpl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: can't load template %s: %s\n", *postTmpl, err)
		os.Exit(1)
	}
	tmplName := filepath.Base(*postTmpl)
	tmpl := template.Must(template.New(tmplName).ParseFiles(tmplPath))
	if err := BuildTree(*dstDir, *srcDir, tmpl); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
