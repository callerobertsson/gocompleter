// Package completer provides the Completer usful for finding matching
// entries in a map of string to interface{}
package completer

import (
	"container/list"
	"errors"
	"strings"
)

// Completer provides function for finding matching completions
type Completer struct {
	m map[string]interface{}
}

// New creates an empty Completer
func New() Completer {
	return Completer{
		m: map[string]interface{}{},
	}
}

// Add adds an entry to the Completer map
func (c Completer) Add(key string, val interface{}) {
	c.m[key] = val
}

// NewFromMap creates a Completer with an input map
func NewFromMap(completerMap map[string]interface{}) Completer {
	return Completer{
		m: completerMap,
	}
}

// Match finds matching entries in the map
func (c Completer) Match(pkey string) (val interface{}, matches []string, err error) {

	matches = c.matchingKeys(pkey)

	if len(matches) != 1 {
		err = errors.New("no unique match found")
		return
	}

	if val, ok := c.m[matches[0]]; ok {
		return val, matches, nil
	}

	err = errors.New("handler for command %q not found")

	return
}

// matchingKeys returns matching strings
// If partial match it returns all matching entries
// If exact match it returns the one matching only.
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
