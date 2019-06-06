package deck

import (
	"sort"
)

// Sorter is an interface used to sort a deck of cards
type Sorter interface {
	Less(deck Deck, i, j int) bool
}

type sortProcessor struct {
	deck   Deck
	sorter Sorter
}

// NewSorter returns a
func NewSorter(sorter Sorter) Processor {
	return sortProcessor{
		sorter: sorter,
	}
}

func (sorterOption sortProcessor) Len() int {
	return sorterOption.deck.Len()
}

func (sorterOption sortProcessor) Less(i, j int) bool {
	return sorterOption.sorter.Less(sorterOption.deck, i, j)
}

func (sorterOption sortProcessor) Swap(i, j int) {
	sorterOption.deck.Swap(i, j)
}

// Process sorts a deck of cards
func (sorterOption sortProcessor) Process(deck Deck) {
	sorterOption.deck = deck
	sort.Sort(sorterOption)
}
