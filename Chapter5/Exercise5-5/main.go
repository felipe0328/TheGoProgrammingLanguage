// Exercise 5.5 Implement CountWordAndImages. (See Exercise 4.9 for word-splitting)
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide an URL to analyze")
		os.Exit(1)
	}

	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Printf("Error counting words and images: %v", err)
	}

	fmt.Printf("Found:\nImages: %d\nWords: %d\n", images, words)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %w", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImages process the html.Node to return the number of words and images in it
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		reader := strings.NewReader(n.Data)
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			value := scanner.Text()
			if value != "" && value != "," && value != " " {
				words++
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	var wordsFC, imagesFC, wordsNS, imagesNS int

	if n.FirstChild != nil {
		wordsFC, imagesFC = countWordsAndImages(n.FirstChild)
	}

	if n.NextSibling != nil {
		wordsNS, imagesNS = countWordsAndImages(n.NextSibling)
	}

	words += wordsFC + wordsNS
	images += imagesFC + imagesNS

	return
}
