// Exercise 5.17 Write a variadic function ElementsByTagName that,
// given an HTML node tree and zero or more names, returns all the elements
// that match one of those names.
package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	exampleHtml := `
		<html>
			<head>
			</head>
			<body>
				<h1>Oscar Puentes</h1>
				<img src="google.com" />
				<h2> End </h2>
			</body>
		</html>
	`

	reader := strings.NewReader(exampleHtml)
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := ElementsByTagName(doc, "h1", "h2")

	for _, element := range result {
		fmt.Printf("%+v\n", element)
	}
}

func ElementsByTagName(n *html.Node, name ...string) []*html.Node {
	if len(name) == 0 {
		return nil
	}

	var resulIndex []*html.Node

	nameIndex := make(map[string]bool)

	for _, value := range name {
		nameIndex[value] = true
	}

	if n.Type == html.ElementNode {
		if nameIndex[n.Data] {
			resulIndex = append(resulIndex, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		resulIndex = append(resulIndex, ElementsByTagName(c, name...)...)
	}

	return resulIndex
}
