package deck

import (
	"fmt"
)

func ExampleNewClone() {
	deck := NewArrayDeck(NewStandard(), NewClone(3))

	fmt.Printf("ExampleNewClone: Deck has %d cards", deck.Len())

	// Output: ExampleNewClone: Deck has 208 cards
}
