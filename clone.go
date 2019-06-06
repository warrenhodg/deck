package deck

type clone struct {
	count int
}

// NewClone returns a processor that clones an existing deck the specified number of times
func NewClone(count int) Processor {
	return clone{count}
}

func (c clone) Process(deck Deck) {
	existing := deck.All()

	for i := 0; i < c.count; i++ {
		deck.AddRange(existing)
	}
}
