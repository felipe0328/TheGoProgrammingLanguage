// Exercise 3.10 Write a non-recursive version of comma, using bytes.Buffer intestead of string concatenation
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567999"))
}

func comma(s string) string {
	var result bytes.Buffer

	for i, val := range s {
		result.WriteRune(val)
		difference := (i + 1) - len(s)
		if difference != 0 && difference%3 == 0 {
			result.WriteRune(',')
		}
	}

	return result.String()
}
