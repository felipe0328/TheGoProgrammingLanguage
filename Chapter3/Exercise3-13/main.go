// Exercise 3.13 Write const declarations for KB, MB, up to YB as compactly as you can
package main

import "fmt"

const (
	KiB = 1000
	MiB = 1000 * KiB
	GiB = 1000 * MiB
)

func main() {
	fmt.Printf("Kb: %d\n", KiB)
	fmt.Printf("Mb: %d\n", MiB)
	fmt.Printf("Gb: %d\n", GiB)
}
