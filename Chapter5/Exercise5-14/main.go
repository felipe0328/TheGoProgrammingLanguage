// Exercise 5.14 Use the breadthFirst function to explore a different structure. For example, you could use the course
// dependencies from the topoSort example (a directed graph), the file system hierarchy on your computer (a tree), or
// a list of bus or subway routes downlaoded from your city government's web site (an undirected graph).
package main

import (
	"fmt"
)

type Tree struct {
	value int
	L     *Tree
	R     *Tree
}

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item Tree) []*Tree, worklist []*Tree) {
	seen := make(map[Tree]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[*item] {
				seen[*item] = true
				worklist = append(worklist, f(*item)...)
			}
		}
	}
}

//!-breadthFirst

// !+crawl
// This crawl inspects the different levels of a binary tree
func crawl(tree Tree) []*Tree {
	fmt.Println(tree.value)

	treeChilds := make([]*Tree, 0)

	if tree.L != nil {
		treeChilds = append(treeChilds, tree.L)
	}

	if tree.R != nil {
		treeChilds = append(treeChilds, tree.R)
	}

	return treeChilds
}

//!-crawl

// !+main
func main() {
	root := &Tree{value: 1}
	root.L = &Tree{value: 2}
	root.R = &Tree{value: 3}
	root.L.L = &Tree{value: 4}
	root.L.R = &Tree{value: 5}
	root.R.L = &Tree{value: 6}

	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, []*Tree{root})
}

//!-main
