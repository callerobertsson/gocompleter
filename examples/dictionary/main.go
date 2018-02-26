// Dictionary Completer example application
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	comp "github.com/callerobertsson/gocompleter"
)

// Path to the dictionary file
var dictPath = "./web2"

// Read dictionary and start interaction
func main() {
	fmt.Printf("Dictionary completer test application\n")

	c, err := createDictionaryComparer(dictPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading dictionary %q: %v\n",
			dictPath, err.Error())
		os.Exit(0)
	}

	fmt.Println("Enter partial search term and return. Ctrl-C to exit")
	interact(c)
}

// createDictionaryComparer() create a ...
func createDictionaryComparer(filePath string) (comp.Completer, error) {

	dictBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return comp.Completer{}, err
	}

	lines := strings.Split(string(dictBytes), "\n")

	c := comp.New()

	for _, line := range lines {
		val := strings.TrimSpace(line)
		if val != "" {
			c.Add(val, true)
		}
	}

	return c, nil
}

// Get input from user and match it using the Completer
func interact(completer comp.Completer) {
	in := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(" > ")
		bs, _, _ := in.ReadLine()
		s := strings.TrimSpace(string(bs))
		_, ms, _ := completer.Match(s)

		limit := 5
		if limit > len(ms) {
			limit = len(ms)
		}

		fmt.Printf("%d of %d matches: %v\n", limit, len(ms), ms[:limit])
	}
}
