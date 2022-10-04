// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	fmt.Println(countWords(src))
}


func countWords(src string) int {
	words := strings.Split(src, " ")
	var newWords []string
	for _, word := range words {
		if word == "" {
			continue
		} else {
			newWords = append(newWords, word)
		}
	}
	return len(newWords)
}


func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")

	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}