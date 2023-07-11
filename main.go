package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	serveDir = flag.String("serveDir", "target", "directory of files to serve")
	port     = flag.Int("port", 8080, "port on which to serve")
)

func main() {
	flag.Parse()

	// TODO: set up logs
	// TODO: set up metrics
	// TODO: set up traces
	// TODO: set up in-memory cache

	log.Println("serving static files from:", *serveDir)
	log.Println("serving on port:", *port)
	http.Handle("/", handler(*serveDir))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Println("fatal:", err)
	}
}
