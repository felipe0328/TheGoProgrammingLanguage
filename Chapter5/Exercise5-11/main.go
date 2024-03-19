// Exercise 5.11 The instructor of linear algebra course decides that calculus is not a prerequisite.
// Extend the topoSort function to report cycles
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

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

// !+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string, parent string)

	visitAll = func(items []string, parent string) {
		for _, item := range items {
			if strings.Contains(parent, item) {
				fmt.Printf("Cyclic Dependency Found! in item %s\n", item)
				os.Exit(1)
				return
			}

			if !seen[item] {
				seen[item] = true
				visitAll(m[item], parent+item)
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys, "")
	return order
}

//!-main
