package main

import "fmt"

// Declaring a custom type
type deck []string

func newDeck() deck {
	cards := deck{}

	var cardSuites = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	var cardValues = []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Jack", "King", "Queen"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// This is a receiver function. Any variable of type deck can call this function like this => cards.print()
// d is the reference to the actual copy of deck type variable, i.e, cards slice in main.go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, " ", card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
