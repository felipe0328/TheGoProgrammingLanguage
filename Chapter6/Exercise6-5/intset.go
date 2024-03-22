package main

import (
	"bytes"
	"fmt"
)

const div = 32 << (^uint(0) >> 63)

// An intset is a set of small non-negative integers.
// Its zero value represents an empty set
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/div, uint(x%div)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/div, uint(x%div)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

// AddAll adds all the elements of the list to the set
func (s *IntSet) AddAll(x ...int) {
	for _, val := range x {
		s.Add(val)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	var counter int
	for _, word := range s.words {
		for j := 0; j < div; j++ {
			if word&(1<<uint(j)) != 0 {
				counter++
			}
		}
	}
	return counter
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}

	word, bit := x/div, uint(x%div)

	if word > len(s.words) { // the word don't exists in the array
		return
	}

	s.words[word] ^= 1 << bit
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

// Copy remove all elements from the set
func (s *IntSet) Copy() *IntSet {
	newS := &IntSet{}

	newS.words = append(newS.words, s.words...)

	return newS
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < div; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", div*i+j)
			}
		}
	}

	buf.WriteByte('}')

	return buf.String()
}

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

// Elements retuns the elemens in the set as a int slice
func (s *IntSet) Elems() []int {
	var result []int
	for i, word := range s.words {
		for j := 0; j < div; j++ {
			if word&uint(1<<j) != 0 {
				result = append(result, div*i+j)
			}
		}
	}
	return result
}
