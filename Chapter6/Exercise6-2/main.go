// Exercise 6.2 Define a variadic (*IntSet).AddAll(..int) method
// that allows a list of values to be added, such as s.AddAll(1,2,3)
package main

import (
	"fmt"
)

// AddAll adds all the elements of the list to the set
func (s *IntSet) AddAll(x ...int) {
	for _, val := range x {
		s.Add(val)
	}
}

func main() {
	var x, y IntSet
	x.AddAll(1, 144, 9)
	x1 := x.Copy()
	fmt.Println("x", &x)
	fmt.Println("x1", x1)
	fmt.Println("x Len", x.Len())
	x.Remove(9)
	x.Remove(2)
	fmt.Println(&x)
	x.Clear()
	fmt.Println(&x)

	fmt.Println("x1", x1)

	y.AddAll(9, 42)
	fmt.Println(&y)

	x.UnionWith(&y)
	fmt.Println(&x)

	fmt.Println(x.Has(9), x.Has(123))
}
