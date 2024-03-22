// Exercise 6.3 (*IntSet).UnionWith computes the union of two sets using |,
// the word-paralel bitwise OR operator. Implement methods for IntersectWith,
// DifferenceWith, and SymetricDifference for the corresponding set operations.
// (The symmetric difference of two set contains the elements present in one set or the other but not both).
package main

import "fmt"

// IntersectWith intersect object with value y
func (s *IntSet) IntersectWith(y *IntSet) {
	lowerLen := len(s.words)

	if len(y.words) < lowerLen {
		lowerLen = len(y.words)
	}

	s.words = s.words[:lowerLen]

	for i, yWord := range y.words {
		if i < len(s.words) {
			s.words[i] &= yWord
		}
	}
}

// DifferenceWith gets the difference between the two sets
func (s *IntSet) DifferenceWith(y *IntSet) {
	for i, yWord := range y.words {
		if i < len(s.words) {
			s.words[i] &^= yWord // s*not(y)
		}
	}
}

// SymetricDifference finds the symectric difference between s and y,
// that means, the numbers that are in s or y, but not in both
func (s *IntSet) SymetricDifference(y *IntSet) {
	for i, yWord := range y.words {
		if i < len(s.words) {
			s.words[i] ^= yWord
		} else {
			s.words = append(s.words, yWord)
		}
	}
}

func main() {
	var x, y IntSet
	x.AddAll(1, 3, 144, 9, 891)
	fmt.Println("x: ", &x)

	y.AddAll(9, 42, 144, 256)
	fmt.Println("y: ", &y)

	x1 := x.Copy()
	x1.IntersectWith(&y)
	fmt.Println("x interesction y: ", x1)

	x2 := x.Copy()
	x2.DifferenceWith(&y)
	fmt.Println("x difference y:", x2)

	x3 := x.Copy()
	x3.SymetricDifference(&y)
	fmt.Println("x difference y:", x3)
}
