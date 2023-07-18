package main

import (
	"expvar"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"testing"
)

type server struct {
	addr     string
	server   *http.Server
	counters counters
}

func serve(t *testing.T, dir string) (*server, error) {
	t.Helper()
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}
	counters := newCounters()
	s := &http.Server{
		Handler: newHandler(dir, counters),
	}
	go s.Serve(l)
	return &server{l.Addr().String(), s, counters}, nil
}

func fetch(t *testing.T, url string) (string, error) {
	t.Helper()
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func expectGet(t *testing.T, url string, want string) {
	if got, err := fetch(t, url); err != nil {
		t.Fatal(err)
	} else if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func expectValue(t *testing.T, counter *expvar.Int, want int64) {
	if counter.Value() != want {
		t.Fatalf("expected counter %s value %d, got %d", counter, want, counter.Value())
	}
}

func TestCaching(t *testing.T) {
	// set up a basic site
	dir := t.TempDir()
	html := `<html>cached?</html>`
	file := path.Join(dir, "cached.html")
	if err := os.WriteFile(file, []byte(html), 0750); err != nil {
		t.Fatal(err)
	}

	// serve it
	s, err := serve(t, dir)
	if err != nil {
		t.Fatal(err)
	}

	// make a test request, expect no cache hit
	url := fmt.Sprintf("http://%s/cached.html", s.addr)
	expectGet(t, url, html)
	expectValue(t, s.counters.cacheHits, 0)

	// make another test request, expect a cache hit
	expectGet(t, url, html)
	expectValue(t, s.counters.cacheHits, 1)

	// make another test request, expect another cache hit
	expectGet(t, url, html)
	expectValue(t, s.counters.cacheHits, 2)

	// shut down
	if err := s.server.Close(); err != nil {
		t.Fatal(err)
	}
}
