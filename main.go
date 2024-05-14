package main

import (
	"flag"
	"log"
)

var (
	shouldBuild = flag.Bool("build", true, "whether to build the site from source")
	shouldServe = flag.Bool("serve", true, "whether to serve the site over http")
	tmplPath    = flag.String("tmpl", "templates/post-template.html", "path to the go html template used to render markdown")
	srcDir      = flag.String("src", "pages", "path to the directory for the site sources")
	siteDir     = flag.String("site", "out", "path to the directory for the built site")
	port        = flag.Int("port", 8080, "port to bind to")
)

func main() {
	flag.Parse()
	if *shouldBuild {
		if err := build(*tmplPath, *srcDir, *siteDir); err != nil {
			log.Fatal(err)
		}
	}
	if *shouldServe {
		if err := serve(*siteDir, *port); err != nil {
			log.Fatal(err)
		}
	}
}
