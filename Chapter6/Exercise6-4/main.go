// Exercise 6.4 Add a method Elems that returns a slice containing the elements of the set,
// suitable for iterating over with a range loop
package main

import "fmt"

// Elements retuns the elemens in the set as a int slice
func (s *IntSet) Elems() []int {
	var result []int
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&uint64(1<<j) != 0 {
				result = append(result, 64*i+j)
			}
		}
	}
	return result
}

func main() {
	var x IntSet
	x.AddAll(1, 99, 5, 128)
	fmt.Println(x.Elems())
}
