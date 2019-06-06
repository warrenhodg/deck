package deck

import (
	"math/rand"
	"time"
)

// Shuffle shuffles a deck of cards
type shuffle struct {
	rng *rand.Rand
}

// NewShuffle returns a new Shuffle object
func NewShuffle(seed int64) Processor {
	return shuffle{
		rng: rand.New(rand.NewSource(seed)),
	}
}

// ShuffleSeedTime returns the current time for use as a random number generator seed
func ShuffleSeedTime() int64 {
	return time.Now().UnixNano()
}

// Process shuffles a deck of cards
func (s shuffle) Process(d Deck) {
	l := d.Len()
	for i1 := 0; i1 < l; i1++ {
		i2 := int(s.rng.Int31n(int32(l)))

		d.Swap(i1, i2)
	}
}
