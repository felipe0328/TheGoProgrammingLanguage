// Exercise 5.1 Change the findlinks program to traverse the n.FirstChild linked list using recursive calls to visit
// instead of a loop
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
		fmt.Println("Please provide a URL for links finding")
		return
	}

	result := fetch(os.Args[1])
	doc, err := html.Parse(result)
	result.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for i, link := range visit(nil, doc) {
		fmt.Println(i, link)
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

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		visit(links, n.NextSibling)
	}

	return links
}
