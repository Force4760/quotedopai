package quotes

import "math/rand"

// go:generate hasgo -T=Quote -S=Quotes
type Quotes []Quote

const NoQuoteString = "Sorry, couldn't find a valid quote!"

func (qs Quotes) Random() Quote {
	length := len(qs)

	// can't choose from an empty quote list
	if length == 0 {
		return NoQuoteString
	}

	// get a random index
	index := rand.Intn(length)

	return qs[index]
}

func (qs Quotes) WordIn(word string) Quote {
	// Filter qs for quotes containing word
	// Return a random one
	return qs.
		Filter(
			func(q Quote) bool {
				return q.InQuote(word)
			},
		).
		Random()
}

func (qs Quotes) WordNotIn(word string) Quote {
	// Filter qs for quotes not containing word
	// Return a random one
	return QuotesDoPai.
		Filter(
			func(q Quote) bool {
				return q.NotInQuote(word)
			},
		).
		Random()
}

func (qs Quotes) WordAndIn(word1, word2 string) Quote {
	// Filter qs for quotes containing word1 && word2
	// Return a random one
	return QuotesDoPai.
		Filter(
			func(q Quote) bool {
				return q.AndInQuote(word1, word2)
			},
		).
		Random()
}

func (qs Quotes) WordOrIn(word1, word2 string) Quote {
	// Filter qs for quotes containing word1 || word2
	// Return a random one
	return QuotesDoPai.
		Filter(
			func(q Quote) bool {
				return q.OrInQuote(word1, word2)
			},
		).
		Random()
}

func (qs Quotes) WordXorIn(word1, word2 string) Quote {
	// Filter qs for quotes containing word1 XOR word2
	// Return a random one
	return QuotesDoPai.
		Filter(
			func(q Quote) bool {
				return q.XorInQuote(word1, word2)
			},
		).
		Random()
}
