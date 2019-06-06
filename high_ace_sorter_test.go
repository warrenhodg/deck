package deck

import (
	"fmt"
)

func ExampleNewHighAceSorter() {
	std := NewStandard()
	shuffle := NewShuffle(fixedShuffleSeed)
	sorter := NewSorter(NewHighAceSorter())
	deck := NewArrayDeck(std, shuffle, sorter)

	fmt.Printf("ExampleNewHighAceSorter: Card 0 is %s", deck.Card(0))

	// Output: ExampleNewHighAceSorter: Card 0 is 2 of spades
}
