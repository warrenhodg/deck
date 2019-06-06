package deck

import (
	"fmt"
)

func ExampleNewNumericalSorter() {
	std := NewStandard()
	shuffle := NewShuffle(fixedShuffleSeed)
	sorter := NewSorter(NewNumericalSorter())
	deck := NewArrayDeck(std, shuffle, sorter)

	fmt.Printf("ExampleNewNumericalSorter: Card 0 is %s", deck.Card(0))

	// Output: ExampleNewNumericalSorter: Card 0 is ace of spades
}
