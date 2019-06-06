package deck

import (
	"fmt"
)

func ExampleNewArrayDeck() {
	deck := NewArrayDeck()

	fmt.Printf("ExampleNew: Deck has %d cards", deck.Len())

	// Output: ExampleNew: Deck has 0 cards
}
