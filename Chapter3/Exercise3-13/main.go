// Exercise 3.13 Write const declarations for KB, MB, up to YB as compactly as you can
package main

import "fmt"

const (
	KiB = 1000
	MiB = KiB * KiB
	GiB = KiB * MiB
	TiB = KiB * GiB
	// and so on..
)

func main() {
	fmt.Printf("Kb: %d\n", KiB)
	fmt.Printf("Mb: %d\n", MiB)
	fmt.Printf("Gb: %d\n", GiB)
}
