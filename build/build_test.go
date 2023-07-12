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

func check(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func expect[U any](t *testing.T) func(U, error) U {
	return func(u U, err error) U {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
		return u
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
	src := t.TempDir()

	html := "<html><body>Foo</body></html>"
	check(t, os.WriteFile(path.Join(src, "a.html"), []byte(html), 0750))

	md := "# hello"
	check(t, os.WriteFile(path.Join(src, "b.md"), []byte(md), 0750))
	check(t, os.Mkdir(path.Join(src, "foo"), 0750))
	check(t, os.WriteFile(path.Join(src, "foo/c.md"), []byte(md), 0750))

	dst := t.TempDir()
	if err := BuildTree(dst, src); err != nil {
		t.Fatal(err)
	}

	var dstPaths []string
	filepath.WalkDir(dst, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			relPath, err := filepath.Rel(dst, path)
			if err != nil {
				t.Fatal(err)
			}
			dstPaths = append(dstPaths, relPath)
		}
		return nil
	})

	want := []string{"a.html", "b.md", "foo/c.md"}
	if !shallowSliceEq(dstPaths, want) {
		t.Fatalf("want paths: %v, got %v", want, dstPaths)
	}

	// TODO: check that the file is actually processed
}
