// Exercise 7.18 Using the tocket-based decoder API, write a program that will read an arbitrary
// XML document and construct a tree of generic nodes that represent it. Nodes are of two
// kinds: CharData nodes represent text strings, and Element nodes represent named elements
// and their attributes. Each element node has a slice of child nodes.

// TODO: This exercise is going to be resolved later
package main

import "encoding/xml"

// You may find the following declaration helpfull

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Name
	Children []Node
}

func main() {}
