package deck

import (
	"os"
"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 52{
		t.Errorf("Expected decks lenght of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Clubs"{
		t.Errorf("Expected Ace of Clubs got %v", d[0])
	}
	if d[len(d)-1] != "King of Spades"{
		t.Errorf("Expected King of Spades got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	deck := NewDeck()

	deck.SaveToFile("_decktesting")

	loadedDeck := NewDeckFromFile("_decktesting")
	if len(loadedDeck) != 52{
		t.Errorf("Expected 52 cards in decks, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}