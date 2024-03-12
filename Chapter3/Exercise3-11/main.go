// Exercise 3.11 Enhance comma so that it deals correctly with floating-point numbers and an optional sign
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(comma(scanner.Text()))
	}
}

func comma(s string) string {
	var result bytes.Buffer
	if s[0] == '-' {
		result.WriteRune('-')
		s = s[1:]
	}

	data := strings.Split(s, ".")

	var floats string
	if len(data) > 1 {
		floats = "." + data[1]
		s = data[0]
	}

	for i, val := range s {
		result.WriteRune(val)
		difference := (i + 1) - len(s)
		if difference != 0 && difference%3 == 0 {
			result.WriteRune(',')
		}
	}

	result.WriteString(floats)

	return result.String()
}
