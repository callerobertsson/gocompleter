package completer

import (
	"container/list"
	"errors"
	"strings"
)

type Completer struct {
	m map[string]interface{}
}

func New(completerMap map[string]interface{}) Completer {
	return Completer{
		m: completerMap,
	}
}

// Finds a MapFunc from a partial matching key
func (c Completer) Get(pkey string) (val interface{}, matches []string, err error) {

	matches = c.matchingKeys(pkey)

	if len(matches) != 1 {
		err = errors.New("No unique match found")
		return
	}

	if val, ok := c.m[matches[0]]; ok {
		return val, matches, nil
	}

	err = errors.New("Handler for command %q not found")

	return
}

func (c Completer) matchingKeys(str string) []string {

	matchList := list.New()

	// Check for exact match
	if _, ok := c.m[str]; ok {
		return []string{str}
	}

	// Check for matches starting with str
	for val := range c.m {
		if strings.HasPrefix(val, str) {
			matchList.PushBack(val)
		}
	}

	matches := make([]string, matchList.Len())

	i := 0
	for e := matchList.Front(); e != nil; e = e.Next() {
		matches[i] = e.Value.(string)
		i++
	}

	return matches
}
