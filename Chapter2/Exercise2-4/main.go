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

	inpt := bufio.NewScanner(os.Stdin)

	for inpt.Scan() {
		number, err := strconv.Atoi(inpt.Text())
		if err != nil {
			fmt.Printf("Unable to convert %s: %v\n", inpt.Text(), err)
			continue
		}

		number8 := uint64(number)
		binary := strconv.FormatInt(int64(number), 2)
		fmt.Println(binary, PopCountShift(number8))
	}

}

func PopCountShift(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}

	return count
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
