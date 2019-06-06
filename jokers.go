package deck

type jokers struct {
	count int
}

// NewJokers returns a processor that adds a specified number of jokers to a deck
func NewJokers(count int) Processor {
	return jokers{
		count: count,
	}
}

// Process adds the specified number of jokers to the deck
func (j jokers) Process(deck Deck) {
	for i := 0; i < j.count; i++ {
		deck.Add(Card{
			Number: BadNumber,
			Suite:  Joker,
		})
	}
}
