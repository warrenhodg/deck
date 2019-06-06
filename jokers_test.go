package deck

import (
	"fmt"
)

func ExampleNewJokers() {
	deck := NewArrayDeck(NewJokers(4))

	fmt.Printf("ExampleNewJokers: Deck has %d cards, and card 0 is a %s", deck.Len(), deck.Card(0))

	// Output: ExampleNewJokers: Deck has 4 cards, and card 0 is a joker
}
