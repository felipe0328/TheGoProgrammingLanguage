// Exercise 7.3 Write a String method for the *tree type in gopl.io/ch4/treesort
// that reveals the sequence of values in the tree
package main

import (
	"fmt"
	"strconv"
)

func (t *tree) String() string {
	var data string

	if t != nil {
		data += t.left.String()
		data += strconv.Itoa(t.value)
		data += t.right.String()
	}

	return data
}

func main() {
	t := add(nil, 2)
	add(t, 1)
	add(t, 5)
	add(t, 3)
	add(t, 8)

	fmt.Println(t)
}
