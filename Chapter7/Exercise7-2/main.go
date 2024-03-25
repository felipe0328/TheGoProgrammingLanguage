// Exercise 7.2 Write a function CountingWriter with the signature below that,
// given an io.Writer, returns a new Writer that wraps the original, and a pointer
// to an int64 variable that at any moment contains the number of bytes written to
// the new Writer.
// func CountingWriter(w io.Writer)(io.Writer, *int64)
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type WriterWrapper struct {
	w io.Writer
	c int64
}

func (w *WriterWrapper) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	w.c += int64(n)

	return
}

// CountingWriter returns a new io.Writer wrapping the original,
// and creating a pointer int64 that count the number of bytes written.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wrapper := &WriterWrapper{w: w}
	return wrapper, &wrapper.c
}

func main() {
	w := bufio.NewWriter(os.Stdout)

	newW, bCounter := CountingWriter(w)

	fmt.Fprint(newW, "Adding data to new Writer, should be seen in the original output.")
	fmt.Println("Writen bytes: ", *bCounter)

	fmt.Fprint(newW, "Second line of added data")
	fmt.Println("Writen bytes: ", *bCounter)
}
