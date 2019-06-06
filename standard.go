package deck

type standard struct{}

// NewStandard returns a processor that adds a standard set of cards to a deck
func NewStandard() Processor {
	return standard{}
}

func (standard) Process(deck Deck) {
	cardNo := 0
	for s := Spades; s <= Diamonds; s++ {
		for n := Ace; n <= King; n++ {
			deck.Add(Card{
				Number: n,
				Suite:  s,
			})

			cardNo++
		}
	}
}
