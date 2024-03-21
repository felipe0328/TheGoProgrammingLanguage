// Exercise 5.18 Without changing the behavior, rewrite the fetch function to use defer to close the writable file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println(fetch(os.Args[1]))
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	file, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		fError := file.Close()
		if err == nil {
			err = fError
		}
	}()

	nBytes, err := io.Copy(file, resp.Body)

	return local, nBytes, err
}
