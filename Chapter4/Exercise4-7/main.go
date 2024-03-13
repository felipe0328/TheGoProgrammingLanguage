// Exercise 4.7 Modify reverse to reverse the characters of a []byte slice that represent a UTF-8-encoded
// string, in place. Can you do it without allocating new memory?
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	inputText := []byte("This is the Ý input data.")
	reverseUTF(inputText)
	fmt.Println(string(inputText))
}

func reverseUTF(b []byte) {

	// first we reverse the nonAscii bytes so when reverting all the slice the unicode object remains correct
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}

	reverse(b)
}

func reverse(b []byte) {
	size := len(b)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
