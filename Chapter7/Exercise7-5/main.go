// Exercise 7.5 The LimitReader function in the io package accepts an io.Reader r and
// a number of bytes n, and returns another Reader that reads from r but reports and
// end-of-file condition after n bytes. Implement it.
// func LimitReader(r io.Reader, n int64) io.Reader
package main

import (
	"fmt"
	"io"
	"strings"
)

type readerLimiter struct {
	r io.Reader
	n int64
}

func (rl *readerLimiter) Read(p []byte) (n int, err error) {
	if int64(len(p)) > rl.n {
		p = p[:rl.n]
	}

	n, err = rl.r.Read(p)
	rl.n -= int64(n)

	if rl.n <= 0 {
		return n, io.EOF
	}

	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	rl := &readerLimiter{r: r, n: n}
	return rl
}

func main() {
	reader := strings.NewReader("This is a test reader")

	newReader := LimitReader(reader, 8)

	b := make([]byte, 10)
	n, err := newReader.Read(b)
	fmt.Println(n, err, string(b))
}
