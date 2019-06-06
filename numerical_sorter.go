package deck

type numericalSorter struct{}

// NewNumericalSorter returns a new Shuffle object
func NewNumericalSorter() Sorter {
	return numericalSorter{}
}

func (sorter numericalSorter) compareSuite(i, j Suite) int {
	switch {
	case i < j:
		return -1
	case i == j:
		return 0
	default:
		return 1
	}
}

func (sorter numericalSorter) compareNumber(i, j Number) int {
	switch {
	case i < j:
		return -1
	case i == j:
		return 0
	default:
		return 1
	}
}

// Less is used for sorting a deck of cards
func (sorter numericalSorter) Less(deck Deck, i, j int) bool {
	c1 := deck.Card(i)
	c2 := deck.Card(j)

	s := sorter.compareSuite(c1.Suite, c2.Suite)
	n := sorter.compareNumber(c1.Number, c2.Number)

	return n < 0 || (n == 0 && s < 1)
}
