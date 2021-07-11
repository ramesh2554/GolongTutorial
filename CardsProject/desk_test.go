package main

import (
	"os"
	"testing"
)

/*
Testing for desk.go
foe testing in go file name should be ends with _test
*/

// testing for new deck
func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of spades but got %v ", d[0])
	}
	if d[len(d)-1] != "Four of club" {
		t.Errorf("Expected Last card of Four of Clubs but got %v ", d[len(d)-1])
	}
}

// testing for save ing a file to HDD

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.RemoveAll("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck , got %v ", len(loadedDeck))

	}
	os.RemoveAll("_decktesting")
}
