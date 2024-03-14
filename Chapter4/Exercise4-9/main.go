// Exercuse 4.9 Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const fileName = "input.txt"

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	defer f.Close()

	words := make(map[string]int, 0)

	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		word = strings.Trim(word, ".,- '´")
		words[word]++
	}

	wordNames := make([]string, 0, len(words))

	for name := range words {
		wordNames = append(wordNames, name)
	}

	sort.Slice(wordNames, func(i, j int) bool {
		return words[wordNames[i]] > words[wordNames[j]]
	})

	fmt.Printf("Word\tFrequency\n")
	for _, word := range wordNames {
		fmt.Printf("%10s\t%2d\n", word, words[word])
	}
}
