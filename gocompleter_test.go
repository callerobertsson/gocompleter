package completer

import (
	"testing"
)

type handlerFunc func(interface{}) (interface{}, error)

var dummyMap = map[string]interface{}{
	"a":   nilHandler,
	"abc": nilHandler,
	"abd": nilHandler,
	"abe": nilHandler,
}

func nilHandler(i interface{}) (interface{}, error) {
	return nil, nil
}

func TestGetHandler(t *testing.T) {

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
			s:   "doesnotexist",
			num: 0,
		},
	}

	fm := New(dummyMap)

	for _, c := range cases {

		_, ms, _ := fm.Get(c.s)

		if len(ms) != c.num {

			t.Errorf("GetHandler(%v) returned %v matches, expected %v",
				c.s, len(ms), c.num)
		}
	}
}