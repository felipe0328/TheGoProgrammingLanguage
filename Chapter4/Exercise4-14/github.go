package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	githubApi   = "https://api.github.com"
	repoFetcher = "/repos/%s/%s/%s"
	issues      = "issues"
	milestones  = "milestones"
	assignees   = "assignees"
)

type fetchType interface {
	[]GitHubIssues | []Milestone | []User
}

type GitHubIssues struct {
	Title   string
	State   string
	User    *User
	Number  int
	HTMLURL string `json:"html_url"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title       string
	HTMLURL     string `json:"html_url"`
	State       string
	Creator     *User
	Number  int
}

func FetchData[T fetchType](owner, repo, tipe string) (*T, error) {
	issuesUrl := fmt.Sprintf(repoFetcher, owner, repo, tipe)
	data := new(T)
	err := fetchFromGit(issuesUrl, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchFromGit[T fetchType](url string, object *T) error {
	githubUrl := fmt.Sprintf("%s%s", githubApi, url)
	resp, err := http.Get(githubUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(object); err != nil {
		return err
	}

	return nil
}
