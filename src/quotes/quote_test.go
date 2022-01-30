package quotes

import "testing"

func TestInQuote(t *testing.T) {
	quote1 := Quote("")
	if quote1.InQuote("ABC") {
		t.Log("Empty Quote shouldn't contain any string")
		t.Fail()
	}

	quote2 := Quote("AB")
	if quote2.InQuote("ABC") {
		t.Log("AB Quote shouldn't contain ABC")
		t.Fail()
	}

	quote3 := Quote("ABC")
	if !quote3.InQuote("ABC") {
		t.Log("ABC Quote should contain ABC")
		t.Fail()
	}

	quote4 := Quote("ABCD")
	if !quote4.InQuote("ABC") {
		t.Log("ABCD Quote should contain ABC")
		t.Fail()
	}

	quote5 := Quote("aBCD")
	if quote5.InQuote("ABC") {
		t.Log("aBCD Quote shouldn't contain ABC")
		t.Fail()
	}
}

func TestNotInQuote(t *testing.T) {
	quote1 := Quote("")
	if !quote1.NotInQuote("ABC") {
		t.Log("Empty Quote shouldn't contain any string")
		t.Fail()
	}

	quote2 := Quote("AB")
	if !quote2.NotInQuote("ABC") {
		t.Log("AB Quote shouldn't contain ABC")
		t.Fail()
	}

	quote3 := Quote("ABC")
	if quote3.NotInQuote("ABC") {
		t.Log("ABC Quote should contain ABC")
		t.Fail()
	}

	quote4 := Quote("ABCD")
	if quote4.NotInQuote("ABC") {
		t.Log("ABCD Quote should contain ABC")
		t.Fail()
	}

	quote5 := Quote("aBCD")
	if !quote5.NotInQuote("ABC") {
		t.Log("aBCD Quote shouldn't contain ABC")
		t.Fail()
	}
}

func TestAndInQuote(t *testing.T) {
	quote1 := Quote("")
	if quote1.AndInQuote("ABC", "DEF") {
		t.Log("Empty Quote shouldn't contain any string")
		t.Fail()
	}

	quote2 := Quote("AB")
	if quote2.AndInQuote("ABC", "DEF") {
		t.Log("AB Quote shouldn't contain neither ABC nor DEF")
		t.Fail()
	}

	quote3 := Quote("ABC")
	if quote3.AndInQuote("ABC", "DEF") {
		t.Log("ABC Quote shouldn't contain DEF")
		t.Fail()
	}

	quote4 := Quote("ABCDEF")
	if !quote4.AndInQuote("ABC", "DEF") {
		t.Log("ABCDEF Quote should contain ABC and DEF")
		t.Fail()
	}

	quote5 := Quote("aBCdEF")
	if quote5.AndInQuote("ABC", "DEF") {
		t.Log("aBCdEF Quote shouldn't contain neither ABC nor DEF")
		t.Fail()
	}
}

func TestOrInQuote(t *testing.T) {
	quote1 := Quote("")
	if quote1.OrInQuote("ABC", "DEF") {
		t.Log("Empty Quote shouldn't contain any string")
		t.Fail()
	}

	quote2 := Quote("AB")
	if quote2.OrInQuote("ABC", "DEF") {
		t.Log("AB Quote shouldn't contain neither ABC nor DEF")
		t.Fail()
	}

	quote3 := Quote("ABC")
	if !quote3.OrInQuote("ABC", "DEF") {
		t.Log("ABC Quote should contain ABC")
		t.Fail()
	}

	quote4 := Quote("ABCDEF")
	if !quote4.OrInQuote("ABC", "DEF") {
		t.Log("ABCDEF Quote should contain ABC and DEF")
		t.Fail()
	}

	quote5 := Quote("aBCdEF")
	if quote5.OrInQuote("ABC", "DEF") {
		t.Log("aBCdEF Quote shouldn't contain neither ABC nor DEF")
		t.Fail()
	}
}

func TestXorInQuote(t *testing.T) {
	quote1 := Quote("")
	if quote1.XorInQuote("ABC", "DEF") {
		t.Log("Empty Quote shouldn't contain any string")
		t.Fail()
	}

	quote2 := Quote("AB")
	if quote2.XorInQuote("ABC", "DEF") {
		t.Log("AB Quote shouldn't contain neither ABC nor DEF")
		t.Fail()
	}

	quote3 := Quote("ABC")
	if !quote3.XorInQuote("ABC", "DEF") {
		t.Log("ABC Quote should contain ABC and not DEF")
		t.Fail()
	}

	quote4 := Quote("ABCDEF")
	if quote4.XorInQuote("ABC", "DEF") {
		t.Log("ABCDEF Quote should contain ABC and DEF")
		t.Fail()
	}

	quote5 := Quote("aBCdEF")
	if quote5.XorInQuote("ABC", "DEF") {
		t.Log("aBCdEF Quote shouldn't contain neither ABC nor DEF")
		t.Fail()
	}
}
