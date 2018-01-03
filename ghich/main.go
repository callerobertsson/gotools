// Package main implements ghich (aka which) command for finding executables in path
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	allFlag     bool
	verboseFlag bool
	helpFlag    bool
)

func main() {
	flag.BoolVar(&allFlag, "a", false, "show all matches")
	flag.BoolVar(&verboseFlag, "v", false, "add for verbose output")
	flag.BoolVar(&helpFlag, "h", false, "show usage info and exit")
	flag.Parse()
	arg := flag.Arg(0)

	if arg == "" || helpFlag {
		usage()
		os.Exit(0)
	}

	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")

	verbose("Looping $PATH searching for executable %q\n", arg)

	which(arg, paths)
}

func which(exe string, paths []string) {

	pad := ""
	if verboseFlag {
		pad = "  "
	}

	for _, path := range paths {
		verbose("%v:\n", path)

		dir, err := os.Open(path)
		if err != nil {
			fmt.Printf("%vcould not open path %v\n", pad, path)
			continue
		}

		fi, err := dir.Stat()
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		if !fi.IsDir() {
			fmt.Printf("%v%v is not a directory\n", pad, path)
			continue
		}

		fis, err := dir.Readdir(0)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		for _, fi := range fis {
			if fi.Name() == exe {
				msg := ""
				if fi.Mode()&0111 == 0 {
					msg = " <-- warning: not executable"
				}
				fmt.Printf("%v%v%c%v%v\n", pad, path, os.PathSeparator, fi.Name(), msg)

				if !allFlag {
					return
				}
				break
			}
		}

	}
}

func verbose(format string, a ...interface{}) {
	if verboseFlag {
		fmt.Printf(format, a...)
	}
}

func usage() {
	fmt.Println("Synopsis: ghich [-h] [-v] [-a] <command_name>")
	flag.PrintDefaults()
}
