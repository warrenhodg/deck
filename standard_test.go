package deck

import (
	"fmt"
)

func ExampleNewStandard() {
	deck := NewArrayDeck(NewStandard())

	fmt.Printf("ExampleNewStandard: Deck has %d cards", deck.Len())

	// Output: ExampleNewStandard: Deck has 52 cards
}
