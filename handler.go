package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type observedResponseWriter struct {
	http.ResponseWriter
	req *http.Request
}

type request struct {
	time      time.Time
	client    string
	userAgent string
	method    string
	host      string
	url       *url.URL
	status    int
}

func (req request) String() string {
	return fmt.Sprintf(
		`%s "%s" %s %s %s %d`,
		req.client, req.userAgent, req.method, req.host, req.url, req.status)
}

func (o observedResponseWriter) WriteHeader(statusCode int) {
	e := request{
		time:      time.Now(),
		client:    o.req.RemoteAddr,
		userAgent: o.req.UserAgent(),
		method:    o.req.Method,
		host:      o.req.Host,
		url:       o.req.URL,
		status:    statusCode,
	}
	log.Println(e)
	o.ResponseWriter.WriteHeader(statusCode)
}

func observe(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		o := observedResponseWriter{res, req}
		h.ServeHTTP(&o, req)
	})
}

func handler(serveDir string) http.Handler {
	fs := http.FileServer(http.Dir(serveDir))
	return observe(fs)
}
