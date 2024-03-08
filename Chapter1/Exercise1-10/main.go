// Exercise 1.10: Find a web site that produces a large amount of data.
// Investigae caching by running fetchall twice in succession to see whether
// the reported time changes match. Do you get the same content each time?.
// Modify fetchall to print its output to a file so it can be examined.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const urlEndpoint = "https://data.wa.gov/api/views/f6w7-q2d2/rows.csv?accessType=DOWNLOAD"

func main() {
	start := time.Now()
	ch := make(chan string)

	go fetch(urlEndpoint, ch, 1)
	go fetch(urlEndpoint, ch, 2)

	// NOTE: Both files are exaclty the same, so, no changes, but when running to goroutines one takes longer than the other
	// also, when running twice the program the secode time is slower than the first time

	counter := 0
	for _, url := range os.Args[1:] {
		counter++
		localCounter := counter
		go fetch(url, ch, localCounter) // starts a goroutine
	}

	for range os.Args[1:] { // we need one more index to get all the data in channel
		fmt.Println(<-ch) // receives data from channel ch
	}

	for i := 0; i < 2; i++ {
		fmt.Println(<-ch) // Adding a reading for the added fetchs
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string, index int) {
	start := time.Now()

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	f, err := os.Create(fmt.Sprintf("file_%d.txt", index))
	if err != nil {
		ch <- fmt.Sprintf("While creating file %s: %v", url, err)
		return
	}
	defer f.Close()

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
