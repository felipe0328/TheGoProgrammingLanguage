// Exercise 4.10 Modify issues to report the results in age categories, say less than a month old, less than a year old
// and more than a year old.
package main

import (
	"Exercise4-10/github"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	Month = time.Hour * 24 * 30
	Year  = Month * 12
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %10.10s\t%-20s\n", item.Number, item.User.Login, item.Title, convertTimeToAge(item.CreatedAt))
	}
}

func convertTimeToAge(t time.Time) string {
	now := time.Now()

	duration := now.Sub(t)

	if duration < Month {
		return "less than a month old"
	} else if duration < Year {
		return "less than a year old"
	}

	return "more than a year old"
}
