package deck

import (
	"fmt"
)

func ExampleNewFilter() {
	deck := NewArrayDeck(NewStandard(), NewFilter(Card{Three, Spades}))

	fmt.Printf("ExampleNewFilter: Deck has %d cards, card 2 is %s", deck.Len(), deck.Card(2))

	// Output: ExampleNewFilter: Deck has 51 cards, card 2 is 4 of spades
}
