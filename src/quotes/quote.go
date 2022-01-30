package quotes

import (
	"strings"
)

type Quote string

func (q Quote) InQuote(word string) bool {
	// check if word is in q
	return strings.Contains(string(q), word)
}

func (q Quote) NotInQuote(word string) bool {
	// check if word is not in q
	return !q.InQuote(word)
}

func (q Quote) AndInQuote(word1, word2 string) bool {
	// check if word1 && word2 are in q
	contains1 := strings.Contains(string(q), word1)
	contains2 := strings.Contains(string(q), word2)

	return contains1 && contains2
}

func (q Quote) OrInQuote(word1, word2 string) bool {
	// check if word1 || word2 are in q
	contains1 := strings.Contains(string(q), word1)
	contains2 := strings.Contains(string(q), word2)

	return contains1 || contains2
}

func (q Quote) XorInQuote(word1, word2 string) bool {
	// check if word1 XOR word2 are in q
	contains1 := strings.Contains(string(q), word1)
	contains2 := strings.Contains(string(q), word2)

	// Go does not have a built-in XOR operator
	// != is somewhat equivalent, sice it has the same properties for bools
	// T F -> T
	// F T -> T
	// T T -> F
	// F F -> F
	return contains1 != contains2
}
