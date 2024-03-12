// Exercise 3.12 Write a function that reports whether two strings are anagrams of each other, that is, they contain
// the same letters in a different order
package main

import (
	"fmt"
	"strings"
)

func main() {

	var s1, s2 string
	fmt.Print("Please insert the first word: ")
	fmt.Scanln(&s1)

	fmt.Print("Please insert the second word: ")
	fmt.Scanln(&s2)

	fmt.Println(anagrams(s1, s2))
}

func anagrams(s1, s2 string) bool {
	s1, s2 = strings.ToLower(s1), strings.ToLower(s2)

	comparer := make(map[rune]int)

	for _, element := range s1 {
		comparer[element]++
	}

	for _, element := range s2 {
		comparer[element]--

		if comparer[element] == 0 {
			delete(comparer, element)
		}
	}

	return len(comparer) == 0
}
