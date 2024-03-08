// Try fetchall with longer arguments lists, such as samples from the top million
// web sites available at alexa.com.
// How does the program behave if a web site just doesn't responde?
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const jsonFile = "ranked_domains.json"

type jsonStruct struct {
	Position int
	Domain   string
	Count    int
	Etv      float32
}

// NOTE: I added a timeout to cancel the requests if the webpage takes to much to finish the request
// if the website didn't response the execution gets freezed until receiving a response.
func main() {
	start := time.Now()
	ch := make(chan string)

	f, err := os.Open(jsonFile)
	if err != nil {
		fmt.Printf("Error while opening file %s: %v", jsonFile, err)
		return
	}

	var data []jsonStruct
	fileData, err := io.ReadAll(f)
	f.Close()

	if err != nil {
		fmt.Printf("Error while reading file %s: %v", jsonFile, err)
		return
	}
	json.Unmarshal(fileData, &data)

	for _, url := range data {
		go fetch(url.Domain, ch) // starts a goroutine
	}

	for range data { // we need one more index to get all the data in channel
		fmt.Println(<-ch) // receives data from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 50*time.Second) // adding a 50 seconds timeout

	req, _ := http.NewRequest("GET", url, nil)
	req = req.WithContext(ctx)
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
