// Package deck implements a Deck and its Cards
package deck

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var colors = []string{"❤", "♦", "♣", "♠"}

// Deck is a full deck of 52 cards
type Deck struct {
	Cards Cards
}

// Card represents a standard playing card
type Card struct {
	Color string
	Value int
}

// Cards is a slice of Card
type Cards []Card

// Hand is just another name for Cards used differently
type Hand Cards

// NewDeck creates a Deck with 52 shuffled cards
func NewDeck() *Deck {
	d := Deck{getCardsForDeck()}
	d.Shuffle()

	return &d
}

// Shuffle mixes all cards in the Deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Draw takes one card from the deck
func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) < 1 {
		return Card{}, fmt.Errorf("no more cards")
	}

	c := d.Cards[0]
	d.Cards = d.Cards[1:]

	return c, nil
}

// String returns the textual representation of a Card
func (c Card) String() string {
	v := strconv.Itoa(c.Value)
	switch c.Value {
	case 13:
		v = "K"
	case 12:
		v = "Q"
	case 11:
		v = "J"
	}

	return fmt.Sprintf("%s%v", c.Color, v)
}

// String returns the textual representation of a Hand of Cards
func (h Hand) String() string {
	s := ""
	for _, c := range h {
		s += fmt.Sprintf(" %v", c)
	}
	return s
}

// CalculatedSum counts the sum of the Cards according to game rules
func (h Hand) CalculatedSum() int {
	sum := 0

	for _, c := range h {
		val := c.Value
		if val == 13 {
			val = 11 // ace
			if sum+11 > 21 {
				val = 1
			}

		} else if val > 10 {
			val = 10 // king, queen, jack
		}
		sum += val
	}

	return sum
}

// Returns 52 cards of different color and value
func getCardsForDeck() Cards {
	cs := Cards{}
	for _, c := range colors {
		for i := 1; i <= 13; i++ {
			cs = append(cs, Card{c, i})
		}
	}

	return cs
}
