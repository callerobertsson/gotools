// package main implements a cli that checks if input is a valid Luhn number
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "Need number as argument\n")
		os.Exit(1)
	}

	answer := "invalid"
	if luhn(flag.Arg(0)) {
		answer = "valid"
	}

	fmt.Printf("%v\n", answer)
}

func luhn(s string) bool {
	l := len(s)
	as := []int{}
	sum := 0
	for i, v := range strings.Split(s, "") {
		a, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if (l-i)%2 == 0 {
			a *= 2
		}
		if a > 9 {
			a -= 9
		}
		sum += a
		as = append(as, a)
	}

	return sum%10 == 0
}
