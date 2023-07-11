package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	srcDir = flag.String("srcDir", "pages", "source directory to build")
	dstDir = flag.String("dstDir", "target", "build output directory")
)

func main() {
	log.Println("source directory:", *srcDir)
	log.Println("output directory:", *dstDir)
	if err := buildTree(*dstDir, *srcDir); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
