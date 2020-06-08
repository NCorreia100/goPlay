package main

import "fmt"

func main() {
	cards := newDeck()
	hand, cardsRemaining := deal(cards, 3)

	saveToFile("hand.txt", hand)
	savedHand := newDeckFromFile("hand.txt")

	fmt.Println("saved hand:", savedHand)

}
