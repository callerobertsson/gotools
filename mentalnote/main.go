// Package main implements mentalnote, a simple and fast mental note logger
package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/callerobertsson/gotools/mentalnote/config"
	"github.com/callerobertsson/gotools/mentalnote/store"
	"github.com/callerobertsson/gotools/mentalnote/view"
)

// Default values
const (
	configFileName   = ".mn.json"
	configFileEnvVar = "MENTALNOTE_CONFIG"
)

// Command line flags and options
var (
	flagConfig  string
	flagMessage string
	optionList  bool
)

func init() {
	// Parse command line
	flag.StringVar(&flagConfig, "c", "", "config file path")
	flag.BoolVar(&optionList, "l", false, "list messages")
	flag.StringVar(&flagMessage, "m", "", "message to send")
	flag.Parse()

	// Set config file, if not on command line
	if flagConfig == "" {
		user, err := user.Current()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get current user: %v\n", err)
			os.Exit(1)
		}

		// Set default config file
		flagConfig = user.HomeDir + "/" + configFileName

		// Overide if env var is set
		if os.Getenv(configFileEnvVar) != "" {
			flagConfig = os.Getenv(configFileEnvVar)
		}
		fmt.Printf("Using config file %q\n", flagConfig)
	}
}

func main() {
	// Create config
	cfg, err := config.New(flagConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		os.Exit(1)
	}

	// Check if -l option
	if optionList {
		view.ListMessages(cfg)
		os.Exit(0)
	}

	// If no -m option, read from console
	if flagMessage == "" {
		flagMessage = view.ReadMessageFromConsole()
	}

	slack.SendMessage(flagMessage, cfg)
	fmt.Println("done!")
	os.Exit(0)

	fmt.Println("nothing sent!")
	os.Exit(1)
}
