// Exercise 5.3 Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into
// <script> or <style> elements, since their contents are not visible in a web browser
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: You need to provide an URL for analysis")
		os.Exit(1)
	}

	data := fetch(os.Args[1])
	webNode, err := html.Parse(data)
	if err != nil {
		fmt.Printf("Error parsing HTML: %v", err)
		os.Exit(1)
	}

	content(webNode, "")
}

func fetch(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting %s: %v", url, err)
		os.Exit(1)
	}

	return resp.Body
}

func content(n *html.Node, indent string) {
	nodeContent := strings.TrimSpace(n.Data)

	dataAtom := n.DataAtom.String()

	if dataAtom != "script" && dataAtom != "style" && nodeContent != "" {
		fmt.Printf("%s%.55s\n", indent, nodeContent)
	}

	if n.FirstChild != nil {
		content(n.FirstChild, " ")
	}

	if n.NextSibling != nil {
		content(n.NextSibling, indent+" ")
	}

}
