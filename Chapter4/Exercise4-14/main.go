// Exercise 4.14 Create a web server that queries GitHub once and then allows navigation of the list of bug reports,
// milestones and users.
package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
)

var (
	owner = flag.String("o", "", "Git project owner")
	repo  = flag.String("r", "", "Git Repo")
)

func main() {
	flag.Parse()

	if *owner == "" || *repo == "" {
		fmt.Println("Please provide a GitHub owner and repo.")
		return
	}

	issues, err := FetchData[[]GitHubIssues](*owner, *repo, issues)
	if err != nil {
		fmt.Println("Unable to fetch issues, ", err)
	}

	milestones, err := FetchData[[]Milestone](*owner, *repo, milestones)
	if err != nil {
		fmt.Println("Unable to fetch issues, ", err)
	}

	users, err := FetchData[[]User](*owner, *repo, assignees)
	if err != nil {
		fmt.Println("Unable to fetch issues, ", err)
	}

	data := make(map[string]interface{})
	data["issues"] = *issues
	data["milestones"] = *milestones
	data["users"] = *users

	report := template.Must(template.New("reportData").Parse(templateToUse))

	if err = report.Execute(os.Stdout, data); err != nil {
		fmt.Println("Error creating template: ", err)
	}

}
