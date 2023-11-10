package serve

import (
	"log"
	"net/http"
	"strconv"
)

type observedResponseWriter struct {
	http.ResponseWriter
	req      *http.Request
	counters *counters
}

func (o observedResponseWriter) WriteHeader(statusCode int) {
	log.Printf(
		`%s "%s" %s %s %s %d`,
		o.req.RemoteAddr, o.req.UserAgent(), o.req.Method, o.req.Host,
		o.req.URL.String(), statusCode)
	o.ResponseWriter.WriteHeader(statusCode)
	o.counters.resps.Add(1)
	o.counters.statusCodes.Add(strconv.Itoa(statusCode), 1)
	// only record pages where load succeeds so caller can't oom us
	if statusCode < 400 {
		o.counters.pages.Add(o.req.URL.String(), 1)
	}
}

type observedHandler struct {
	h        http.Handler
	counters counters
}

const cachePolicy = "public, max-age=3600"

func (s *observedHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.counters.reqs.Add(1)
	res.Header().Add("Cache-Control", cachePolicy)
	o := observedResponseWriter{res, req, &s.counters}
	s.h.ServeHTTP(&o, req)
}

func NewHandler(serveDir string) http.Handler {
	fs := http.FileServer(http.Dir(serveDir))
	counters := newCounters()
	return &observedHandler{h: fs, counters: counters}
}
