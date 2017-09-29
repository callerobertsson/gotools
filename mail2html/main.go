// Package main implements a raw mail to HTML converter.
// It extracts the text/html part, if present, and prints it.
// Otherwise it returns the text/plain part with newlines translated to <br/>.
// Suitable to use in conjunction with mutt and w3m.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jhillyerd/enmime"
)

func main() {
	e, err := enmime.ReadEnvelope(os.Stdin)
	if err != nil {
		exit(err)
	}

	if e.HTML != "" {
		fmt.Println(e.HTML)
		exit(nil)
	}

	if e.Text != "" {
		fmt.Println(strings.Replace(e.Text, "\n", "<br/>", -1))
		exit(nil)
	}

	exit(fmt.Errorf("could not find HTML or text parts in mail with subject %q", e.Root.Header.Get("Subject")))
}

func exit(err error) {
	if err == nil {
		os.Exit(0)
	}

	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
