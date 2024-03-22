package main

import (
	"bytes"
	"fmt"
)

// An intset is a set of small non-negative integers.
// Its zero value represents an empty set
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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
		for j := 0; j < 64; j++ {
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

	word, bit := x/64, uint(x%64)

	if word > len(s.words) { // the word don't exists in the array
		return
	}

	s.words[word] ^= 1 << bit
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	buf.WriteByte('}')

	return buf.String()
}