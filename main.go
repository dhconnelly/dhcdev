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

	// TODO: set up logging
	// TODO: set up metrics
	// TODO: set up traces
	// TODO: set up health check
	// TODO: set up in-memory cache
	// TODO: set up proxy around fileserver to log errors

	log.Println("serving static files from:", *serveDir)
	log.Println("serving on port:", *port)
	fs := http.FileServer(http.Dir(*serveDir))
	http.Handle("/", fs)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
