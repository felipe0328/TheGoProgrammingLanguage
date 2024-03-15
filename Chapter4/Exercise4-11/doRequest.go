package main

import "net/http"

func doRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
