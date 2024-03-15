// Exercise 4.11 Build a tool that lets users create, read, update and close GitHub issues from the command line,
// invoking their preferred text editor when substantial text input is required.
package main

import (
	"flag"
	"fmt"
	"strings"
)

type IssueRequest struct {
	Endpoint, Method string
}

type IssueGet struct {
	Title string
	Body  string
}

var (
	CreateIssue = IssueRequest{"https://api.github.com/repos/%s/%s/issues", "POST"}
	GetIssue    = IssueRequest{"https://api.github.com/repos/%s/%s/issues/%d", "GET"}
	UpdateIssue = IssueRequest{"https://api.github.com/repos/%s/%s/issues", "Patch"}
)

var (
	owner  = flag.String("ow", "", "Owner of the repository")
	repo   = flag.String("r", "", "Repository")
	issue  = flag.Int("i", 0, "Issue to search")
	action = flag.String("a", "", "Action, can be get, update, close, create")
)

func main() {
	flag.Parse()

	actionP := strings.ToLower(*action)

	switch actionP {
	case "get":
		fmt.Println(getIssue(*owner, *repo, *issue))
	case "create":
		fmt.Println(createIssue(*owner, *repo))
	case "update":
		updateIssue()
	default:
		fmt.Println("Please provide an action to execute")
	}
}
