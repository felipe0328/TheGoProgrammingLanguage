// Exercise 4.5 Write a in-place function to eliminate adjacent duplicates in a []string slice
package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 3, 4, 5, 5, 6}
	fmt.Println(deleteAdjacentDuplicates(data))
	fmt.Println(data)
}

func deleteAdjacentDuplicates(s []int) []int {
	out := s[:0]
	for i, s1 := range s {
		if i == 0 || s1 != s[i-1] {
			out = append(out, s1)
		}
	}

	return out
}
