// Package rawrunereader implements a single character input helper.
package rawrunereader

import (
	"bufio"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// RawRuneReader is a raw rune input reader.
type RawRuneReader struct {
	in    *bufio.Reader
	state *terminal.State
}

// New creates a new RawRuneReader instance.
func New() (*RawRuneReader, error) {

	state, err := terminal.MakeRaw(syscall.Stdin)
	if err != nil {
		return nil, err
	}

	rcr := RawRuneReader{}
	rcr.state = state
	rcr.in = bufio.NewReader(os.Stdin)

	return &rcr, nil
}

// ReadRune reads a character (rune) from STDIN.
func (rcr *RawRuneReader) ReadRune() (rune, error) {
	r, _, err := rcr.in.ReadRune()
	// fmt.Printf("rune: %#v\r\n", r)
	return r, err
}

// Restore restores the terminal to its previous state.
// Best is to call it in a defer after the creation of the RawRuneReader.
func (rcr *RawRuneReader) Restore() error {
	return terminal.Restore(syscall.Stdin, rcr.state)
}
