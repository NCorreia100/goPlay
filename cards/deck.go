package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	// var deckArray = []string{}

	deckByte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	deckArray := strings.Split(string(deckByte), ",")

	return deck(deckArray)
}

func (thisD deck) shuffle() {
	randomFactor := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range thisD {
		newPosition := randomFactor.Intn(len(thisD) - 1)
		thisD[i], thisD[newPosition] = thisD[newPosition], thisD[i]
	}
}
