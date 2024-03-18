// Exercise 5.4 Extend the visit function so that it extracts other kinsd of links from the document, such as images
// scripts, and style sheets.
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
		fmt.Println("Please provide an URL for analysis.")
		os.Exit(1)
	}

	data := fetch(os.Args[1])
	defer data.Close()

	webData, err := html.Parse(data)
	if err != nil {
		fmt.Println("Error parsing html: ", err)
		os.Exit(1)
	}

	fmt.Println(visit(nil, webData))
}

func fetch(url string) io.ReadCloser {
	data, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting URL: ", err)
		os.Exit(1)
	}

	return data.Body
}

var infoTypes = map[string]string{
	"a":      "href",
	"script": "src",
	"img":    "src",
	"meta":   "content",
	"link":   "href",
	"style":  "src",
}

func visit(data []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		infoKey, ok := infoTypes[n.Data]
		if ok {
			for _, value := range n.Attr {
				if value.Key == infoKey {
					data = append(data, fmt.Sprintf("%s %s:%s\n", n.Data, value.Key, value.Val))
				}
			}
		}
	}

	if n.FirstChild != nil {
		data = visit(data, n.FirstChild)
	}

	if n.NextSibling != nil {
		data = visit(data, n.NextSibling)
	}

	return data
}
