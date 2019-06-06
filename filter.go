package deck

type filter struct {
	Card
}

// NewFilter returns a processor that removes all of specific type of card
func NewFilter(card Card) Processor {
	return filter{card}
}

// Process adds the specified number of jokers to the deck
func (f filter) Process(deck Deck) {
	i := 0
	for {
		if i == deck.Len() {
			return
		}

		if deck.Card(i).Equals(f.Card) {
			deck.Remove(i)
		} else {
			i++
		}
	}
}
