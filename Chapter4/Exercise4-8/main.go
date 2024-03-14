// Exercise 4.8 Modify charcount to count letters, digits, and so on in ther Unicode categories, using functions
// like unicode.IsLetter
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	letters := make(map[rune]int)   // counts of Unicode letters
	digits := make(map[rune]int)    // counts of Unicode digits
	punct := make(map[rune]int)     // counts of punct elements
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, err
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsLetter(r):
			letters[r]++
		case unicode.IsDigit(r):
			digits[r]++
		case unicode.IsPunct(r):
			punct[r]++
		default:
			counts[r]++
		}

		utflen[n]++
	}
	fmt.Printf("letter\tcount\n")
	for c, n := range letters {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("punct\tcount\n")
	for c, n := range punct {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("digits\tcount\n")
	for c, n := range digits {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("others\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters \n", invalid)
	}
}
