// Package view implements view functions for mentalnote
package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/callerobertsson/gotools/mentalnote/config"
	"github.com/callerobertsson/gotools/mentalnote/store"
)

// ListMessages prints a list of messages to the console
func ListMessages(conf config.Config) {

	messages, err := slack.GetMessages(conf)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	lastDay := ""
	for i := len(messages) - 1; i >= 0; i-- {
		mess := messages[i]
		currDay := mess.GetDateString()

		if lastDay != currDay {
			lastDay = currDay
			fmt.Println("\n" + currDay)
		}
		//text := strings.Replace(mess.Text, "\n", "       \n", -1)
		fmt.Println("  " + mess.GetTimeString() + ": " + mess.Text)
	}
}

// ReadMessageFromConsole builds a message from console input
func ReadMessageFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		if line == "." {
			break
		}

		lines = append(lines, line)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}
