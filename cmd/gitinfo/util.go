package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func exe() string   { return filepath.Base(os.Args[0]) }
func ok()           { exit(0) }
func exit(code int) { os.Exit(code) }

func fail(code int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	exit(code)
} // fail()

func usage(long bool) {
	fmt.Printf("%s [options]\n", exe())
	fmt.Printf("%s [options] <path>\n", exe())

	// do we need to display long usage?
	if long {
		fmt.Println()
		flag.PrintDefaults()
	}

	// we are done
	ok()
} // usage()
