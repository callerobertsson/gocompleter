// Package completer unit tests
package completer

import (
	"testing"
)

var dummyMap = map[string]interface{}{
	"a":   nilHandler,
	"abc": nilHandler,
	"abd": nilHandler,
	"abe": nilHandler,
}

func nilHandler(i interface{}) (interface{}, error) {
	return nil, nil
}

func TestMatchHandler(t *testing.T) {

	cases := []struct {
		s   string
		num int
	}{
		{
			s:   "a",
			num: 1,
		},
		{
			s:   "ab",
			num: 3,
		},
		{
			s:   "",
			num: 4,
		},
		{
			s:   "doesnotexist",
			num: 0,
		},
	}

	fm := NewFromMap(dummyMap)

	for _, c := range cases {

		_, ms, _ := fm.Match(c.s)

		if len(ms) != c.num {

			t.Errorf("MatchHandler(%v) returned %v matches, expected %v",
				c.s, len(ms), c.num)
		}
	}
}

func TestAddHandler(t *testing.T) {

	m := New()
	m.Add("a", "first a")
	m.Add("b", "only b")
	m.Add("a", "a overridden")

	_, ms, _ := m.Match("")

	if len(ms) != 2 {
		t.Errorf("Expected 2 matches but got %d", len(ms))
	}
}
