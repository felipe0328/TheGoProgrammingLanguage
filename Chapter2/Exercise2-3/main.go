// Exercise 2.3 Rewrite PopCount to use a loop instead of a single expression.
// Compare performance of the two versions.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error converting %s: %v\n", scanner.Text(), err)
		}

		input := int64(value)
		fmt.Printf("Input %s contains %d set bits\n", strconv.FormatInt(input, 2), PopCount(uint64(input)))
		fmt.Printf("Input %s contains %d set bits\n", strconv.FormatInt(input, 2), PopCountLoop(uint64(input)))
		fmt.Printf("Input %s contains %d set bits\n", strconv.FormatInt(input, 2), PopCountLoopComplete(uint64(input)))
		fmt.Printf("/=============/\n\n")
	}

}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var result int

	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}

	return result
}

func PopCountLoopComplete(x uint64) int {
	var result int

	for i := 0; i < 8; i++ {
		byteValue := byte(x >> (i * 8))
		result += int(getByteValue(byteValue))
	}

	return result
}

func getByteValue(b byte) byte {
	if b <= 0 {
		return 0
	}

	return getByteValue(b/2) + byte(b&1)
}
