package main

import (
	"fmt"
	"github.com/LoyalEnv0y/CardsOld/deck"
)

func main() {
	cards := deck.NewDeck() // we are calling the deck of cards
	cards.Print() // and printing them one by one
	fmt.Printf("\n")

	// "hand" is the cards we dealt and "remainingCards" are what is left.
	// also we specify the amount of cards that we want to deal
	hand, remainingCards := deck.Deal(cards, 5)

	// we are printing the 5 cards we dealt
	hand.Print()
	// and we are printing what is left of the cards
	remainingCards.Print()

	cards.SaveToFile("my_cards")

	/*	cards := newDeckFromFile("my_cards")
		cards.print()*/

}