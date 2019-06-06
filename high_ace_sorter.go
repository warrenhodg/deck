package deck

// HighAceSorter sorts a deck of cards by HighAce value
type highAceSorter struct{}

// NewHighAceSorter returns a new Shuffle object
func NewHighAceSorter() Sorter {
	return highAceSorter{}
}

func (sorter highAceSorter) compareSuite(i, j Suite) int {
	switch {
	case i < j:
		return -1
	case i == j:
		return 0
	default:
		return 1
	}
}

func (sorter highAceSorter) compareNumber(i, j Number) int {
	switch {
	case i == Ace && j != Ace:
		return 1
	case j == Ace && i != Ace:
		return -1
	case i < j:
		return -1
	case i == j:
		return 0
	default:
		return 1
	}
}

// Less is used for sorting a deck of cards
func (sorter highAceSorter) Less(deck Deck, i, j int) bool {
	c1 := deck.Card(i)
	c2 := deck.Card(j)

	s := sorter.compareSuite(c1.Suite, c2.Suite)
	n := sorter.compareNumber(c1.Number, c2.Number)

	return n < 0 || (n == 0 && s < 1)
}
