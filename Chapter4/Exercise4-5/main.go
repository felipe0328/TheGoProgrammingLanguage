// Exercise 4.5 Write a in-place function to eliminate adjacent duplicates in a []string slice
package main

import "fmt"

func main() {
	data := []string{"a", "b", "b", "data", "data", "other", "data", "final"}
	fmt.Println(deleteAdjacentDuplicates(data))
	fmt.Println(data)
}

func deleteAdjacentDuplicates(s []string) []string {
	out := s[:0]
	for i, s1 := range s {
		if i == 0 || s1 != s[i-1] {
			out = append(out, s1)
		}
	}

	return out
}
