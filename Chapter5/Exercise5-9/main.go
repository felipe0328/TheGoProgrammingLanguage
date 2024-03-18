// Exercise 5.9 Write a function expand(s string, f func(string)string) string that replaces each substring "$foo"
// within s by the test returned by f("foo")
package main

import (
	"fmt"
	"strings"
)

const substring = "$foo"

func main() {
	testInput := "this is an $foo input found in the value $foo"

	fmt.Println(expand(testInput, encoder))
}

func expand(s string, f func(string) string) string {
	values := strings.Split(s, " ")

	var result []string

	for _, value := range values {
		if value == substring {
			value = f(value)
		}
		result = append(result, value)
	}

	return strings.Join(result, " ")
}

func encoder(s string) string {
	final := ""

	for _, val := range s {
		final += string(val + 2)
	}

	return final
}
