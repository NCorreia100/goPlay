package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

//create a new type of "deck" which is a slice of strings
func newDeck() deck {
	cardTypes := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	newDeck := deck{}

	for i := range cardTypes {
		for j := range cardValues {
			newDeck = append(newDeck, cardValues[j]+" of "+cardTypes[i])
		}
	}
	return newDeck
}

func (thisD deck) print() {
	for i, card := range thisD {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	removedCards := d[:handSize]
	remainingDeck := d[handSize:]
	return removedCards, remainingDeck
}
func (thisD deck) toString() string {
	deckSlice := []string(thisD)
	return strings.Join(deckSlice, ",")
}
func saveToFile(filename string, d deck) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	var deckArray = []string{}
	var d = deck{}
	deckStr := string(ioutil.ReadFile(filename))
	deckArray := strings.Split(deckStr, ",")

	for _, card := range deckArray {

	}
}
