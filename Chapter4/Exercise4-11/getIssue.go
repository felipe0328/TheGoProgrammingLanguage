package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssue(owner, repo string, issue int) *IssueGet {
	req, err := http.NewRequest(GetIssue.Method, fmt.Sprintf(GetIssue.Endpoint, owner, repo, issue), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp, err := doRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	var result IssueGet
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	return &result
}
