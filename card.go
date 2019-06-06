package deck

import (
	"fmt"
)

// Number is the number on the card
type Number byte

// Card numbers
const (
	BadNumber = Number(0)
	Ace       = Number(iota)
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Suite represents a card Suite or a joker
type Suite byte

// Card suites
const (
	BadSuite = Suite(0)
	Spades   = Suite(iota)
	Hearts
	Clubs
	Diamonds
	Joker
)

// Card represents a single card in a deck of cards
type Card struct {
	Number Number
	Suite  Suite
}

func (s Suite) String() string {
	switch s {
	case Spades:
		return "spades"
	case Hearts:
		return "hearts"
	case Clubs:
		return "clubs"
	case Diamonds:
		return "diamonds"
	case Joker:
		return "joker"
	default:
		return "?"
	}
}

func (n Number) String() string {
	switch {
	case n == Ace:
		return "ace"
	case n >= Two && n <= Ten:
		return fmt.Sprintf("%d", int(n-Two)+2)
	case n == Jack:
		return "jack"
	case n == Queen:
		return "queen"
	case n == King:
		return "king"
	default:
		return "?"
	}
}

// Equals tests two cards for equality
func (c Card) Equals(other Card) bool {
	switch {
	case c.Number != BadNumber && other.Number != BadNumber && c.Number != other.Number:
		return false

	case c.Suite != BadSuite && other.Suite != BadSuite && c.Suite != other.Suite:
		return false
	default:
		return true
	}
}

func (c Card) String() string {
	switch c.Suite {
	case Joker:
		return c.Suite.String()
	default:
		return fmt.Sprintf("%s of %s", c.Number.String(), c.Suite.String())
	}
}
