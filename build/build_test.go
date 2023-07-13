package main

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func shallowSliceEq[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

func collectPaths(t *testing.T, dir string) ([]string, error) {
	t.Helper()
	var dstPaths []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			dstPaths = append(dstPaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dstPaths, nil
}

func check(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildTreeInvalid(t *testing.T) {
	if err := BuildTree(t.TempDir(), "blah"); err == nil {
		t.Fatal("expected error on nonexistent source directory")
	}
}

func TestBuildTreeEmpty(t *testing.T) {
	dst := t.TempDir()
	src := t.TempDir()
	if err := BuildTree(dst, src); err != nil {
		t.Fatal(err)
	}
	filepath.WalkDir(dst, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			t.Fatal(err)
		}
		if path != dst {
			t.Fatal("no files expected!")
		}
		return nil
	})
}

func TestBuildTree(t *testing.T) {
	// create the site source
	src := t.TempDir()
	html := "<html><body>Foo</body></html>"
	check(t, os.WriteFile(path.Join(src, "a.html"), []byte(html), 0750))
	md := "# hello"
	check(t, os.WriteFile(path.Join(src, "b.md"), []byte(md), 0750))
	check(t, os.Mkdir(path.Join(src, "foo"), 0750))
	check(t, os.WriteFile(path.Join(src, "foo/c.md"), []byte(md), 0750))

	// build the site
	dst := t.TempDir()
	if err := BuildTree(dst, src); err != nil {
		t.Fatal(err)
	}

	// verify the files were all built in the right places
	dstPaths, err := collectPaths(t, dst)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"a.html", "b.md", "foo/c.md"}
	for i := range want {
		want[i] = path.Join(dst, want[i])
	}
	if !shallowSliceEq(dstPaths, want) {
		t.Fatalf("want paths: %v, got %v", want, dstPaths)
	}

	// only verify that the files are non-empty: build logic is validated
	// in a different case
	for _, path := range want {
		if bs, err := os.ReadFile(path); err != nil {
			t.Fatal(err)
		} else if len(bs) == 0 {
			t.Fatalf("built file is empty: %s", path)
		}
	}
}

func TestBuildFileHTML(t *testing.T) {
	dir := t.TempDir()
	srcPath, dstPath := path.Join(dir, "src.html"), path.Join(dir, "dest.html")
	html := "<html><body>foo</body></html>"
	if err := os.WriteFile(srcPath, []byte(html), 0750); err != nil {
		t.Fatal(err)
	}
	src, err := os.Open(srcPath)
	if err != nil {
		t.Fatal(err)
	}
	dst, err := os.Create(dstPath)
	if err != nil {
		t.Fatal(err)
	}
	if err := BuildFile(dst, src); err != nil {
		t.Fatal(err)
	}
	check(t, src.Close())
	check(t, dst.Close())
	got, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(got); got != html {
		t.Fatalf("want %s, got %s", html, got)
	}
}
