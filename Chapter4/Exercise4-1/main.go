// Exercise 4.1 Write a function that counts the number of bits that are different in two SHA256 hashes.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(HashComparer(c1, c2))
}

func HashComparer(c1, c2 [32]byte) int {
	counter := 0

	for i := 0; i < len(c1); i++ {
		bitToCompare := c1[i] ^ c2[i] // We use XOR to get the bits that are different
		counter += getSetBits(bitToCompare)
	}

	return counter
}

func getSetBits(b byte) int {
	count := 0

	mask := uint8(1)

	for b > 0 {
		if b&mask == 1 {
			count++
		}
		b >>= 1
	}

	return count
}
