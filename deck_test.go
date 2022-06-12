package deck

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// First Test
	expected := 16

	if len(d) != expected {
		t.Errorf("Expected deck length of %v, but got %v", expected, len(d))
	}

	// Second Test
	first_card := d[0]
	expected_card := "Ace of Spades"

	if first_card != expected_card {
		t.Errorf("Expected first card is %v, but got %v", expected_card, first_card)
	}

	// Thirst Test
	last_card := d[expected-1]
	expected_card = "Four of Clubs"

	if last_card != expected_card {
		t.Errorf("Expected last card is %v, but got %v", expected_card, last_card)
	}
}
