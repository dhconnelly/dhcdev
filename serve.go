package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func serveDirectory(dir string, port int) error {
	// TODO:
	// https://gowebexamples.com/
	// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
	// https://blog.cloudflare.com/exposing-go-on-the-internet
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("serving files from %s at http://%s", dir, addr)
	fs := http.FS(os.DirFS(dir))
	handler := http.FileServer(fs)
	server := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server.ListenAndServe()
}
