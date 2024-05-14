package main

import (
	"flag"
	"log"
)

var (
	shouldBuild = flag.Bool("build", true, "whether to build the site from source")
	shouldServe = flag.Bool("serve", true, "whether to serve the site over http")
	outDir      = flag.String("out", "out", "path to the directory for the built site")
	srcDir      = flag.String("src", "pages", "path to the directory for the site sources")
	tmplPath    = flag.String("tmpl", "templates/post-template.html", "path to the go html template used to render markdown")
	port        = flag.Int("port", 8080, "port to bind to")
)

func build(outDir, srcDir, tmplPath string) {
	if err := buildSite(outDir, srcDir, tmplPath); err != nil {
		log.Fatalf("build: %s", err)
	}
}

func serve(dir string, port int) {
	if err := serveDirectory(dir, port); err != nil {
		log.Fatalf("serve: %s", err)
	}
}

func main() {
	flag.Parse()
	if *shouldBuild {
		build(*outDir, *srcDir, *tmplPath)
	}
	if *shouldServe {
		serve(*outDir, *port)
	}
}
