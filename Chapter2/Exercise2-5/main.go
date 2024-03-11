package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Error converting %s: %v\n", input, err)
			continue
		}

		fmt.Println(strconv.FormatInt(int64(number), 2), PopCountClearing(uint64(number)))
	}
}

func PopCountClearing(x uint64) int {
	var count int

	for i := 0; i < 64; i++ {
		if x-(x&(x-1)) > 0 {
			count++
		}
		x = x & (x - 1)
	}

	return count
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
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
