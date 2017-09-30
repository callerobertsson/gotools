// Package main implements the gopastebin tool that saves text from stdin to pastebin.com
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// URL to the pastebin.com API
const apiURL = "http://pastebin.com/api/api_post.php"

// Command line options
var (
	k string // user key
	o string // option
	p int    // private
	n string // paste name
	x string // expire
	f string // format
)

func init() {
	// Parse command line options
	flag.StringVar(&k, "k", "", "the api key")
	flag.StringVar(&o, "o", "paste", "api option (paste)")
	flag.IntVar(&p, "p", 1, "0 = public, 1 = unlisted, 2 = private")
	flag.StringVar(&n, "n", "", "paste name")
	flag.StringVar(&x, "x", "1Y", "expiration")
	flag.StringVar(&f, "f", "text", "paste format")
	flag.Parse()
}

func main() {

	// Check for API key
	if k == "" {
		fmt.Fprintf(os.Stderr, "No API key! Please get one from pastebin.com and set the -k option\n")
		os.Exit(1)
	}

	// Read text from stdin
	code, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read from STDIN: %v\n", err)
		os.Exit(1)
	}

	// Call pastebin API
	resp, err := http.PostForm(apiURL, url.Values{
		"api_dev_key":           {k},
		"api_paste_private":     {strconv.Itoa(p)},
		"api_option":            {o},
		"api_paste_name":        {n},
		"api_paste_expire_date": {x},
		"api_paste_format":      {f},
		"api_paste_code":        {string(code)},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
		os.Exit(1)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}

	// Print response
	fmt.Println(string(body))
}
