package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	//if length of deck frof newDeck func != 52 throw error
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected first card as King of Hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//remove possible decktesting file
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	//if length of loadedDeck saved as _deckTesting != 52 throw error
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}
	//remove decktesting file
	os.Remove("_decktesting")
}
