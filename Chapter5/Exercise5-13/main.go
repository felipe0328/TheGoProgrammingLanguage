// Exercise 5.13 Modify crawl to make local copies of the pages it finds, creating directories as necessary.
// Don't make copies of pages that come from a different domain. For example, if the original page comes
// from golang.org, save all files from there, but exlude ones from vimeo.com
package main

import (
	"Exercise5-13/links"
	"fmt"
	"log"
	"os"
	"strings"
)

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

// !+crawl
func crawl(url string) []string {
	if strings.Contains(os.Args[1], url) {
		// Create folder with data
		fmt.Println("Creating folder")
	}
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

// !+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
