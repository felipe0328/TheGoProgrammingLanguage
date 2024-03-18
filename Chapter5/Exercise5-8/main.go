// Exercise 5.8 Modify forEachNode so that the pre and post functions return a boolean result
// indicating whether to continue the transversal. Use it to write a function ElementByID with
// the following signature that finds the first HTML element with the specified id attribute.
// The function should stop the transversal as soon as a match is found
// func ElementByID(doc *html.Node, id string) *html.Node
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("please provide a valid input")
		os.Exit(1)
	}

	doc, err := fetcher(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	node := ElementByID(doc, os.Args[2])
	fmt.Printf("%+v\n", node)
}

func fetcher(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s: %w", url, err)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to parse document url %s: %w", url, err)
	}

	return doc, nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, nil)
}

func forEachNode(doc *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	var isFound bool
	if pre != nil {
		isFound = pre(doc, id)
		if isFound {
			return doc
		}
	}

	if !isFound {
		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			result := forEachNode(c, id, pre, post)
			if result != nil {
				return result
			}
		}

		if post != nil {
			post(doc, id)
		}
	}

	return nil
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, value := range n.Attr {
			if value.Key == "id" {
				if value.Val == id {
					return true
				}
			}
		}
	}

	return false
}
