package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := createDeck() // it create 12 cards
	if len((d)) != 12 {
		t.Errorf("Expected deck length of 12 but got %v", len((d)))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card of Spades of Ace. but got %v", d[0])
	}

	if d[(len(d)-1)] != "Hearts of Four" {
		t.Errorf("Expected last card of Hearts of Four. but got %v", d[(len(d)-1)])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("test.txt")
	deck := createDeck()
	deck.saveToFile("test.txt")

	loadedDeck := readDeckFromFile("test.txt")
	if len(loadedDeck) != 12 {
		t.Errorf("Expected 12 cards in deck, got %v", len(loadedDeck))
	}
}
