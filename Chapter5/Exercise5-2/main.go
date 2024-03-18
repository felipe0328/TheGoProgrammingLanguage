// Exercise 5.2 Write a function to populate a mapping from element names -p, div, span, and so on- to the number of
// elements with that name in an HTML document tree.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL to analize.")
		os.Exit(1)
	}

	resultMap := make(map[string]int)
	htmlData := fetch(os.Args[1])
	parsed, err := html.Parse(htmlData)
	htmlData.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	analizeData(resultMap, parsed)

	for key, value := range resultMap {
		fmt.Printf("%-10s: %5d\n", key, value)
	}
}

func analizeData(resultMap map[string]int, n *html.Node) {
	dataAtom := n.DataAtom.String()

	if dataAtom != "" {
		resultMap[dataAtom]++
	}

	if n.FirstChild != nil {
		analizeData(resultMap, n.FirstChild)
	}

	if n.NextSibling != nil {
		analizeData(resultMap, n.NextSibling)
	}
}

func fetch(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting web, %v\n", err)
		os.Exit(1)
	}

	return resp.Body
}
