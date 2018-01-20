// Package main implements chkpath for checking consistency of $PATH
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	helpFlag    bool
	verboseFlag bool
)

func main() {
	flag.BoolVar(&helpFlag, "h", false, "show usage info and exit")
	flag.BoolVar(&verboseFlag, "v", false, "print verbose info")
	flag.Parse()

	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")

	if helpFlag {
		usage()
		os.Exit(0)
	}

	checkPaths(paths)
}

func checkPaths(paths []string) {

	done := map[string]bool{}

	for _, path := range paths {

		if _, ok := done[path]; ok {
			out(path, "duplicate")
			continue
		}

		done[path] = true

		// Open path
		dir, err := os.Open(path)
		if err != nil {
			out(path, "could not open path")
			continue
		}

		// Get file info
		fi, err := dir.Stat()
		if err != nil {
			out(path, err.Error())
			continue
		}

		// Check if it's a directory
		if !fi.IsDir() {
			out(path, "not a directory")
			continue
		}

		// Try to read directory
		fis, err := dir.Readdir(0)
		if err != nil {
			out("%v", err.Error())
			continue
		}

		// Check if directory contains any executables
		hasExecutable := false
		for _, fi := range fis {
			if fi.Mode()&0111 != 0 {
				hasExecutable = true
				continue
			}
		}

		// Print message if no executables in dir
		if !hasExecutable {
			out(path, "contains no executables")
			continue
		}

		// Print ok if verbose mode
		if verboseFlag {
			out(path, "ok")
		}
	}
}

// out prints path and a message
func out(p string, m string) {
	fmt.Printf("%v - %v\n", p, m)
}

// usage prints usage information
func usage() {
	fmt.Println("Synopsis: chkpath [-h] [-v]")
	flag.PrintDefaults()
}
