// Exercise 4.2 Write a program that prints the SHA256 hash of its standart input by default but support a
// command line flag to print the SHA384 or SHA512 hash instead
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

type hashOptions interface {
	[32]byte | [48]byte | [64]byte
}

type T1 int

func main() {
	const inputValue = "Felipe"

	hashType := flag.Int("hash", 256, "has type, can be 256, 384 or 512")
	flag.Parse()

	switch *hashType {
	case 384:
		fmt.Printf("%x\n", hasher(inputValue, sha512.Sum384))
		return
	case 512:
		fmt.Printf("%x\n", hasher(inputValue, sha512.Sum512))
		return
	default:
		fmt.Printf("%x\n", hasher(inputValue, sha256.Sum256))
		return
	}
}

func hasher[T hashOptions](s string, method func([]byte) T) T {
	return method([]byte(s))
}
