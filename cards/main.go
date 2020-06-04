package main

func main() {
	cards := newDeck()
	hand, cardsRemaining := deal(cards, 3)

	cardsRemaining.print()
	hand.print()
}
