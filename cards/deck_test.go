package main

import (
	"fmt"
	"os"
	"testing"
)

func TestDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length to be 54 but got %v", len(d))
	}

	if d[1] != "2 of Spades" {
		t.Errorf("Expected second element to be 2 of Spades but got %v", d[1])
	}
}

func TestIOForDeck(t *testing.T) {
	err := os.Remove("_testingDeck")
	if err != nil {
		fmt.Println("Unable to remove file, reason:", err, "This is expected.")
	}
	d := newDeck()
	saveToFile("_testingDeck", d)
	newD := newDeckFromFile("_testingDeck")

	if d[1] != newD[1] {
		t.Errorf("Expected saved deck to contain %v but got %v", d[1], newD[1])
	}
	err = os.Remove("_testingDeck")
	if err != nil {
		t.Errorf("Unable to remove file")

	}
}
