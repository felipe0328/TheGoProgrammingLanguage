// Exercise 5.15 Write variadic functions max and min, analogous to sum. What should these functions do when called
// with no arguments?. Write variants that require at least one argument
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(min())
	fmt.Println(min(2, 5, 7, 9, 1))
	fmt.Println(max())
	fmt.Println(max(2, 5, 7, 9, 1))
	fmt.Println(maxV(3, 7, 99, 100))
	fmt.Println(minV(3, 7, 99, 100))
}

func min(n ...int) int {
	if len(n) == 0 {
		return 0
	}

	min := math.MaxInt

	for _, value := range n {
		if value < min {
			min = value
		}
	}

	return min
}

func max(n ...int) int {
	if len(n) == 0 {
		return 0
	}

	max := math.MinInt

	for _, value := range n {
		if value > max {
			max = value
		}
	}

	return max
}

func maxV(n1 int, n ...int) int {
	max := n1

	for _, value := range n {
		if value > max {
			max = value
		}
	}

	return max
}

func minV(n1 int, n ...int) int {
	min := n1

	for _, value := range n {
		if value < min {
			min = value
		}
	}

	return min
}
