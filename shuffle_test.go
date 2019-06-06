package deck

import (
	"fmt"
)

const fixedShuffleSeed = 0xDEADBEEF

func ExampleNewShuffle() {
	std := NewStandard()
	shuffle := NewShuffle(fixedShuffleSeed)
	deck := NewArrayDeck(std, shuffle)

	fmt.Printf("ExampleNewShuffle: Card 0 is %s", deck.Card(0))

	// Output: ExampleNewShuffle: Card 0 is 3 of spades
}
