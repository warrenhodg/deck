package deck

// Deck represents a deck of cards
type Deck interface {
	Len() int
	Swap(i, j int)
	Card(i int) Card
	All() []Card
	Add(card Card)
	AddRange(card []Card)
	Remove(i int) Card
}

// Deck represents an array of cards
type arrayDeck struct {
	cards []Card
}

// Processor adjusts a deck of cards
type Processor interface {
	Process(Deck)
}

// NewArrayDeck returns a new deck of cards
func NewArrayDeck(options ...Processor) Deck {
	cards := make([]Card, 0)

	result := &arrayDeck{
		cards: cards,
	}

	for _, option := range options {
		option.Process(result)
	}

	return result
}

// Len returns the number of cards in the deck
func (deck *arrayDeck) Len() int {
	return len((*deck).cards)
}

// Swap is used to swap two cards in a deck
func (deck *arrayDeck) Swap(i, j int) {
	(*deck).cards[i], (*deck).cards[j] = (*deck).cards[j], (*deck).cards[i]
}

func (deck *arrayDeck) Card(i int) Card {
	if i >= (*deck).Len() {
		return Card{
			Number: BadNumber,
			Suite:  BadSuite,
		}
	}

	return (*deck).cards[i]
}

func (deck *arrayDeck) Add(c Card) {
	(*deck).cards = append((*deck).cards, c)
}

func (deck *arrayDeck) AddRange(c []Card) {
	(*deck).cards = append((*deck).cards, c...)
}

func (deck *arrayDeck) Remove(i int) Card {
	c := deck.Card(i)
	deck.cards = append(deck.cards[:i], deck.cards[i+1:]...)
	return c
}

func (deck *arrayDeck) All() []Card {
	result := make([]Card, deck.Len())
	copy(deck.cards, result)
	return result
}
