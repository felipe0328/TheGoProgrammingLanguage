// Exercise 5.16 Write a variadic version of strings.Join
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(Join(",", "s1", "s2", "s3"))
}

func Join(separator string, s ...string) string {
	result := bytes.NewBuffer([]byte{})

	for i := 0; i < len(s); i++ {
		result.WriteString(s[i])
		if i < len(s)-1 {
			result.WriteString(separator)
		}
	}

	return result.String()
}
