package main

import "fmt"

func main() {
	cards := newDeck() // we are calling the deck of cards
	cards.print() // and printing them one by one
	fmt.Printf("\n")

	// "hand" is the cards we dealt and "remainingCards" are what is left.
	// also we specify the amount of cards that we want to deal
	hand, remainingCards := deal(cards, 5)

	// we are printing the 5 cards we dealt
	hand.print()
	// and we are printing what is left of the cards
	remainingCards.print()

	cards.saveToFile("my_cards")

	/*	cards := newDeckFromFile("my_cards")
		cards.print()*/

}