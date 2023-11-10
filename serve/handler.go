package serve

import (
	"expvar"
	"log"
	"net/http"
	"strconv"
)

var (
	reqs        = expvar.NewInt("requests")
	resps       = expvar.NewInt("responses")
	pages       = expvar.NewMap("pages")
	statusCodes = expvar.NewMap("statusCodes")
)

type observedResponseWriter struct {
	http.ResponseWriter
	req *http.Request
}

func (o observedResponseWriter) WriteHeader(statusCode int) {
	log.Printf(
		`%s "%s" %s %s %s %d`,
		o.req.RemoteAddr, o.req.UserAgent(), o.req.Method, o.req.Host,
		o.req.URL.String(), statusCode)
	o.ResponseWriter.WriteHeader(statusCode)
	resps.Add(1)
	statusCodes.Add(strconv.Itoa(statusCode), 1)

	// only record pages where load succeeds so caller can't oom us
	if statusCode == 200 {
		pages.Add(o.req.URL.String(), 1)
	}
}

type observedHandler struct {
	h http.Handler
}

const cachePolicy = "public, max-age=3600"

func (s *observedHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	reqs.Add(1)
	res.Header().Add("Cache-Control", cachePolicy)
	o := observedResponseWriter{res, req}
	if req.Method == "GET" {
		s.h.ServeHTTP(&o, req)
	} else {
		o.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func NewHandler(root http.FileSystem) http.Handler {
	fs := http.FileServer(root)
	return &observedHandler{h: fs}
}
