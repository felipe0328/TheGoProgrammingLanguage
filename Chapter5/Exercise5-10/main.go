// Exercise 5.10 Rewrite topoSort to use maps instead of slices and eliminate tehe inital sort.
// Verify that the results, though nondeterministic, are valid topological orderings
package main

import (
	"fmt"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(data []string)

	visitAll = func(data []string) {
		for _, item := range data {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for key, value := range m {
		if !seen[key] {
			seen[key] = true
			visitAll(value)
			order = append(order, key)
		}
	}

	return order
}
