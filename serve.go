package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type observedResponseWriter struct {
	http.ResponseWriter
	status int
}

func (ow *observedResponseWriter) WriteHeader(status int) {
	ow.status = status
	ow.ResponseWriter.WriteHeader(status)
}

func observe(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ow := observedResponseWriter{ResponseWriter: w}
		h.ServeHTTP(&ow, r)
		log.Printf(
			"%s %s %s %s %s %d",
			r.RemoteAddr, r.UserAgent(), r.Method, r.Host, r.URL.String(),
			ow.status,
		)
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func router(dir string) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.FileServer(http.Dir(dir)))
	mux.Handle("GET /healthz", http.HandlerFunc(healthCheck))
	return mux
}

func newServer(dir string) http.Handler {
	var handler http.Handler
	handler = router(dir)
	handler = observe(handler)
	return handler
}

func ServeDirectory(dir string, port int) error {
	// https://gowebexamples.com/
	// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
	// https://blog.cloudflare.com/exposing-go-on-the-internet
	// TODO:
	// https://cloud.google.com/stackdriver/docs/instrumentation/setup/go
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("serving files from %s at http://%s", dir, addr)
	server := &http.Server{
		Addr:         addr,
		Handler:      newServer(dir),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}
	return server.ListenAndServe()
}
