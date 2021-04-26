// Package main golang implementation of a Black Jack game
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/callerobertsson/jackblack/deck"
	"github.com/callerobertsson/jackblack/rawrunereader"
)

var gameDeck *deck.Deck
var dealerHand deck.Hand
var playerHand deck.Hand

func main() {
	rr, err := rawrunereader.New()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer rr.Restore()

	for {
		gameDeck = deck.NewDeck()
		playerHand = deck.Hand{}
		dealerHand = deck.Hand{}

		loop(rr)

		fmt.Printf("One more round? [y/N]\r\n")
		r, _ := rr.ReadRune()
		if r != rune('y') {
			fmt.Printf("Thank you for choosing Jack Black!\r\n")
			break
		}
	}

}

func loop(rr *rawrunereader.RawRuneReader) {
	fmt.Println("Jack Black\r")
	fmt.Println("  d = draw (player draws one card)\r")
	fmt.Println("  h = hold (dealer draws)\r")
	fmt.Println("  c = continue (dealer draws until finish)\r")
	fmt.Println("  q = quit game\r")

	forceHold := false

	for {
		r := rune('h')

		if !forceHold {
			tmp, err := rr.ReadRune()
			if err != nil {
				log.Printf("Error reading char: %v\r\n", err)
				break
			}
			r = tmp
		}

		switch r {
		case 'q':
			fmt.Printf("Bye!\r\n")
			return
		case 'd':
			fmt.Printf("-> Draw!\r\n")
			draw()
		case 'h':
			fmt.Printf("-> Hold!\r\n")
			hold()
		case 'c':
			fmt.Printf("-> Continue!\r\n")
			forceHold = true
		}

		renderBoard()

		if hasWinner() {
			break
		}
	}

}

func hasWinner() bool {
	dSum := dealerHand.CalculatedSum()
	pSum := playerHand.CalculatedSum()

	switch {
	case pSum > 21:
		fmt.Printf("=> Player busted!\r\n")
	case dSum > 21:
		fmt.Printf("=> Dealer busted!\r\n")
	case pSum == 21:
		fmt.Printf("=> Player won!\r\n")
	case dSum == 21:
		fmt.Printf("=> Dealer got 21!\r\n")
	case dSum > pSum:
		fmt.Printf("=> Dealer won by score!\r\n")
	case pSum == 21:
		fmt.Printf("Player got 21!\r\n")
	default:
		return false
	}

	return true
}

func draw() {
	c, err := gameDeck.Draw()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	playerHand = append(playerHand, c)
}

func hold() {
	c, err := gameDeck.Draw()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	dealerHand = append(dealerHand, c)
}

func renderBoard() {
	fmt.Printf("Player %2d: %v\r\n", playerHand.CalculatedSum(), playerHand)
	fmt.Printf("Dealer %2d: %v\r\n", dealerHand.CalculatedSum(), dealerHand)
}
