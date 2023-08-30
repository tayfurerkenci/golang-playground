package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// deck is receiver
// by convention the receiver variable is the first/two letter of the type
func (this deck) print() {
	for i, card := range this {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}