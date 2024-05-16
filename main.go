package main

import (
	"flag"
	"fmt"
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

func main() {
	flag.Parse()
	if *shouldBuild {
		if err := BuildSite(*outDir, *srcDir, *tmplPath); err != nil {
			log.Fatalf("build: %s", err)
		}
	}
	if *shouldServe {
		addr := fmt.Sprintf("0.0.0.0:%d", *port)
		if err := ServeDirectory(*outDir, addr); err != nil {
			log.Fatalf("serve: %s", err)
		}
	}
}
