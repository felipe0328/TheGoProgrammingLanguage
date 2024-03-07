// Exercise1.2 Modify the echo program to print the index an value of each of its arguments, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Printf("arg %d: %s\n", index, value)
	}
}
