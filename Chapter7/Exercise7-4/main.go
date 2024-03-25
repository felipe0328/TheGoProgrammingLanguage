// Exercise 7.4 The strings.NewReader function returns a value that satisfies the io.Reader
// interface (and others) by reading from its argument, a string. Implement a simple version
// of NewReader yourself, and use it to make the HTML parser take input from string.
package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type reader string

func (r reader) Read(p []byte) (n int, err error) {

	var counter int
	for i := 0; i < len(p); i++ {
		if len(r) == 0 {
			return counter, io.EOF
		}

		counter++
		p[i] = r[0]
		r = r[1:]
	}

	return counter, nil
}

func main() {
	r := reader("<html><head><title>Test</title></head><body><h1>Title</h1></body></html>")
	data, err := html.Parse(r)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("%+v\n", data)

	visit(data)
}

func visit(n *html.Node) {
	fmt.Printf("%s\n", n.Data)

	if n.FirstChild != nil {
		visit(n.FirstChild)
	}

	if n.NextSibling != nil {
		visit(n.NextSibling)
	}
}
