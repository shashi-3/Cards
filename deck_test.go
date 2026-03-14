package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newdeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of size 52,but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)
	deck := newdeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck of size 52, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
