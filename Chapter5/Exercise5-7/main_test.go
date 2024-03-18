package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestFetchAndGetNodes(t *testing.T) {
	htmlObject := `<html>
			<head>
				<link src='style.com' /> 
			</head>
			<body>
				<h1 class='display=flex'> 
					Felipe Test
				</h1>
			</body>
		</html>`

	expectedOutput := []string{
		"<html>",
		"  <head>",
		"    <link src='style.com' />",
		"  </head>",
		"  <body>",
		"    <h1 class='display=flex'>",
		"      Felipe Test",
		"    </h1>",
		"  </body>",
		"</html>",
	}

	parsed, err := html.Parse(strings.NewReader(htmlObject))

	if err != nil {
		t.Error(err)
	}

	output := &bytes.Buffer{}
	forEachNode(parsed, startElement, endElement, output)

	scanner := bufio.NewScanner(output)
	scanner.Split(bufio.ScanLines)

	counter := 0
	for scanner.Scan() {
		if expectedOutput[counter] != scanner.Text() {
			t.Errorf("Value \n%s\nis different from expected \n%s", scanner.Text(), expectedOutput[counter])
		}
		counter++
	}
}
