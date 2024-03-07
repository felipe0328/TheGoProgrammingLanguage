// Modify dup2 to print the names of all files in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Reading from custom input")
		counts["customInput"] = make(map[string]int)
		countLines(os.Stdin, counts["customInput"])
	} else {
		for _, arg := range files {
			counts[arg] = make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts[arg])
			f.Close()
		}
	}

	for key, value := range counts {
		for line, n := range value {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, key)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	fmt.Println("Getting Data")
	for input.Scan() {
		counts[input.Text()]++
	}

	// Ignoring errors from input.Err()
}
