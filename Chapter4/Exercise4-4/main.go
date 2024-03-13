// Exercise 4.4 Write a version of rotate that operates in a single pass
package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	rotate(4, data)
	fmt.Println(data)

	data2 := []int{1, 2, 3, 4, 5, 6}
	rotateSinglePass(4, data2)
	fmt.Println(data2)
}

func rotateSinglePass(rotation int, s []int) {
	s1 := s[rotation:]
	s1 = append(s1, s[:rotation]...)

	copy(s, s1)
}

func rotate(rotation int, s []int) {
	reverse(s[:rotation])
	reverse(s[rotation:])
	reverse(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
