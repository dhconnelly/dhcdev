package serve

import (
	"bytes"
	"expvar"
	"flag"
	"log"
	"net/http"

	"github.com/dhconnelly/dhcdev/cache"
)

var (
	cacheSize = flag.Int("cacheSize", 10*1000*1000 /* 10 MB */, "cache size in bytes")

	// misses aren't interesting: all 404s are cache misses.
	// instead compare cacheHits with vars from the observedHandler.
	cacheHits = expvar.NewInt("cacheHits")
)

type cachedHandler struct {
	cache cache.Cache
	h     http.Handler
}

func (h *cachedHandler) serveCached(
	resp http.ResponseWriter, req *http.Request, b []byte,
) {
	_, err := resp.Write(b)
	if err != nil {
		log.Printf("failed to write cached response for %s: %s", req.URL.Path, err)
	}
}

type cachedResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (w cachedResponseWriter) Write(b []byte) (int, error) {
	w.buf.Write(b)
	return w.ResponseWriter.Write(b)
}

func (h *cachedHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if data := h.cache.Get(req.URL.Path); data != nil {
		// TODO: this should still update the observed counters
		cacheHits.Add(1)
		h.serveCached(resp, req, data)
		return
	}

	w := cachedResponseWriter{ResponseWriter: resp, buf: &bytes.Buffer{}}
	h.h.ServeHTTP(w, req)
	h.cache.Put(req.URL.Path, w.buf.Bytes())
}

func newCachedHandler(h http.Handler) http.Handler {
	return &cachedHandler{cache: cache.New(*cacheSize), h: h}
}
