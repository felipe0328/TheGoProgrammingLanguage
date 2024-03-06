// Exercise1.3 Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.
// to run the benchmarks use go test -bench=. -benchmem
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(joinsImplementation(os.Args[1:]))
	fmt.Printf("strings.Joins implementation took: %d\n", time.Since(now))

	now = time.Now()
	fmt.Println(customImplementation(os.Args[1:]))
	fmt.Printf("custom implementation took: %d\n", time.Since(now))

	now = time.Now()
	fmt.Println(customImplementationWithBytes(os.Args[1:]))
	fmt.Printf("custom implementation with bytes took: %d\n", time.Since(now))
}

func joinsImplementation(args []string) string {
	return strings.Join(args, " ")
}

func customImplementation(args []string) string {
	var s, sep string
	for _, value := range args {
		s += sep + value
		sep = " "
	}
	return s
}

func customImplementationWithBytes(args []string) string {
	var s bytes.Buffer
	for _, value := range args {
		s.WriteString(value)
		s.WriteString(" ")
	}
	return s.String()
}
