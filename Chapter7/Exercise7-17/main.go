// Exercise 7.17 Extend xmlselect so that elements may be selected not just by name,
// but by their attributes too, in the manner of CSS, so that, for instance, an element
// like <div id="page" class="wide"> could be selected by a matching id or class as well
// as its name.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	err := xmlselect(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
}

func xmlselect(r io.Reader) error {
	dec := xml.NewDecoder(r)
	var stack []string // stack of element names
	counter := 1

	shouldVerifyIdOrClass := false
	inputJoined := strings.Join(os.Args[1:], " ")

	found := false

	if strings.Contains(inputJoined, "id") || strings.Contains(inputJoined, "class") {
		shouldVerifyIdOrClass = true
	}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("xmlselect error : %w", err)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			for _, value := range tok.Attr {
				if shouldVerifyIdOrClass && strings.Contains("idclass", value.Name.Local) {
					stack = append(stack, fmt.Sprintf("%s=%s", value.Name.Local, value.Value))
					counter++
				}
			}
		case xml.EndElement:
			if found {
				fmt.Println("REMOVING")
				found = false
			}

			stack = stack[:len(stack)-counter] // pop
			counter = 1
		case xml.CharData:
			printIfContains(stack, os.Args[1:], string(tok))
		}
	}

	return nil
}

func printIfContains(stack, args []string, tok string) {
	if containsAll(stack, args) {
		fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
