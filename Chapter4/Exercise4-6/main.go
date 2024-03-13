// Exercise 4.6 Write a in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace) in a
// UTF-8-encoded []byte slice into a single ASCII space
package main

import (
	"fmt"
	"unicode"
)

const asciiSpace = byte(32)

func main() {
	inputText := "This is a code full of spaces \t where \n we \f can \vfind \rdifferent\n"
	fmt.Println(inputText)
	fmt.Println(squashSpaces(inputText))
	fmt.Println(inputText)
}

func squashSpaces(s string) string {
	sRune := []rune(s)
	counter := 0
	isCurrentSpace := false
	for _, s1 := range sRune {
		if !unicode.IsSpace(s1) {
			sRune[counter] = s1
			counter++
			isCurrentSpace = false
		} else {
			if isCurrentSpace {
				continue
			}
			isCurrentSpace = true
			sRune[counter] = rune(asciiSpace)
			counter++
		}
	}

	return string(sRune[:counter])
}
