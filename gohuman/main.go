// Package main implements gohuman that converts a number to human readable form
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/callerobertsson/gotools/gohuman/human"
)

var (
	longFlag   bool
	baseOption int
	decOption  int
	numArg     int64
)

// init reads the command line arguments
func init() {
	flag.BoolVar(&longFlag, "l", false, "Long format")
	flag.IntVar(&baseOption, "b", 1024, "Base, 1024 or 1000")
	flag.IntVar(&decOption, "d", 2, "Number of decimals")
	helpFlag := flag.Bool("h", false, "Show help and exit")

	flag.Parse()

	// Show help maybe
	if *helpFlag {
		usage()
		os.Exit(0)
	}

	// Get the number argument as string
	arg := flag.Arg(0)

	// Number argument is mandatory
	if arg == "" {
		usage()
		os.Exit(1)
	}

	// Convert number argument to string
	maybeNumArg, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR! You must specify a number. Error: %v\n", err.Error())
		os.Exit(2)
	}
	numArg = maybeNumArg

}

// main is the primus motor
func main() {
	val := human.Bytes(numArg, decOption, longFlag)

	fmt.Printf("%v\n", val)
}

// usage prints usage information
func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n\n")
	fmt.Fprintf(os.Stderr, "    %s [-l] [-b (1024|1000)] [-d <DECIMALS>] <NUMBER>\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Flags:\n\n")
	flag.PrintDefaults()
}
