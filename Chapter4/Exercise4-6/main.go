// Exercise 4.6 Write a in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace) in a
// UTF-8-encoded []byte slice into a single ASCII space
package main

import (
	"fmt"
	"unicode"
)

const asciiSpace = byte(32)

func main() {
	inputText := []rune("This is a code full of spa   ces \t where \n we \f can \vfind \rdifferent\n examples  .")
	fmt.Println(string(inputText))
	fmt.Println(string(squashSpaces(inputText)))
	fmt.Println(string(inputText))
}

func squashSpaces(s []rune) []rune {
	counter := 0
	isCurrentSpace := false
	for _, s1 := range s {
		if !unicode.IsSpace(rune(s1)) {
			s[counter] = s1
			counter++
			isCurrentSpace = false
		} else {
			if isCurrentSpace {
				continue
			}
			isCurrentSpace = true
			s[counter] = rune(asciiSpace)
			counter++
		}
	}

	return s[:counter]
}
