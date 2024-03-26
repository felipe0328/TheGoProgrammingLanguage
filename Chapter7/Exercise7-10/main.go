// Exercise 7.10 The sort.Interface type can be adapted to other uses. Write a function
// IsPalindrome(s sort.Interface) bool  that reports wheter the sequence s is a palindrome,
// in other words, reversing the sequence would not change it. Assume that the elements at
// indice i and j are equal if !s.Less(i,j) && !s.Less(j,i)
package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	sort.Sort(s)
	isSorted := sort.IsSorted(s)

	sort.Sort(sort.Reverse(s))
	inverseStillSorted := sort.IsSorted(s)

	return isSorted == inverseStillSorted
}

func main() {
	values := []int{1, 1, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(values)))

	values = []int{1, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(values)))
}
