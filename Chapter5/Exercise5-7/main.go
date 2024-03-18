// Exercise 5.7 Develop startElement and endElement into a general HTML pretty-printer.
// Print comment nodes, text nodes, and the attributes of each element (<a href='...'>).
// Use short forms like <img/> instead of <img></img> when an element has no children.
// Write a test to ensure that the output can be parsed suvvesfully.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide a valid url to fetch")
	}

	doc, err := fetchAndGetNodes(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to fetchAndGetNodes %v", err)
	}

	forEachNode(doc, startElement, endElement, os.Stdout)
}

func fetchAndGetNodes(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to get url %s: %w", url, err)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to parse document: %w", err)
	}

	return doc, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, output io.Writer), output io.Writer) {
	if pre != nil {
		pre(n, output)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, output)
	}

	if post != nil {
		post(n, output)
	}
}

var depth int

func startElement(n *html.Node, output io.Writer) {
	end := ">"

	if n.FirstChild == nil {
		end = " />"
	}

	if n.Type == html.ElementNode {
		result := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
		for _, value := range n.Attr {
			result += fmt.Sprintf(" %s='%s'", value.Key, value.Val)
		}

		fmt.Fprintf(output, "%s%s\n", result, end)
		depth++
	}

	if n.Type == html.TextNode {
		data := strings.TrimSpace(n.Data)
		if data != "" {
			fmt.Fprintf(output, "%*s%.50s\n", depth*2, "", data)
		}
	}

}

func endElement(n *html.Node, output io.Writer) {
	if n.Type == html.ElementNode {
		depth--

		if n.FirstChild != nil {
			fmt.Fprintf(output, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
