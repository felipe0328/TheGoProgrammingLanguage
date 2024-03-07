// Exercise 1.9 Modify fetch to also print the HTTP status code, found in resp.Status
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) && !strings.HasPrefix(url, httpsPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "Status Code: %d\n", resp.StatusCode)
	}
}