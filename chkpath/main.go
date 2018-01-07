// Package main implements chkpath for checking consistency of $PATH
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	helpFlag bool
)

func main() {
	flag.BoolVar(&helpFlag, "h", false, "show usage info and exit")
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

	for _, path := range paths {

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

		if !hasExecutable {
			out(path, "contains no executables")
		}
	}
}

func out(p string, m string) {
	fmt.Printf("%v - %v\n", p, m)
}

func usage() {
	fmt.Println("Synopsis: chkpath [-h] [-v] [-a] <command_name>")
	flag.PrintDefaults()
}
