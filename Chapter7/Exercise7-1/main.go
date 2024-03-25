// Exercise 7.1 Using the ideas from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {

	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	var counter int

	for scanner.Scan() {
		*w++
		counter++
	}

	return counter, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	var counter int

	for scanner.Scan() {
		*l++
		counter++
	}

	return counter, nil
}

func main() {
	var b ByteCounter
	var w WordCounter
	var l LineCounter

	testString := "Hello World."

	b.Write([]byte(testString))
	w.Write([]byte(testString))
	l.Write([]byte(testString))

	formatedString := "Welcome \n To Go's \n World."

	fmt.Fprint(&b, formatedString)
	fmt.Fprint(&w, formatedString)
	fmt.Fprint(&l, formatedString)

	fmt.Println(b, w, l)
}
