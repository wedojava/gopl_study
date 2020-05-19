package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	FmtResult(result)
}

func FmtResult(r *github.IssuesSearchResult) {
	now := time.Now()
	format := "#%-5d %9.9s %.55s\n"

	oneMonth := now.AddDate(0, -1, 0)
	oneYear := now.AddDate(-1, 0, 0)

	rsLessMonth := make([]*github.Issue, 0)
	rsLessYear := make([]*github.Issue, 0)
	rsGreaterYear := make([]*github.Issue, 0)

	for _, i := range r.Items {
		switch {
		case i.CreatedAt.After(oneMonth) || i.CreatedAt.Equal(oneMonth):
			rsLessMonth = append(rsLessMonth, i)
		case i.CreatedAt.Before(oneMonth) && i.CreatedAt.After(oneYear):
			rsLessYear = append(rsLessYear, i)
		case i.CreatedAt.Before(oneYear):
			rsGreaterYear = append(rsGreaterYear, i)
		}
	}

	fmt.Println("[!] Issues less than a month:\n-----------------------------")
	if len(rsLessMonth) > 0 {
		for _, i := range rsLessMonth {
			fmt.Printf(format, i.Number, i.User.Login, i.Title)
		}
	}
	fmt.Println("[!] Issuses less than a year:\n-----------------------------")
	if len(rsLessYear) > 0 {
		for _, i := range rsLessYear {
			fmt.Printf(format, i.Number, i.User.Login, i.Title)
		}
	}
	fmt.Println("[!] Issues more than a year:\n-----------------------------")
	if len(rsGreaterYear) > 0 {
		for _, i := range rsGreaterYear {
			fmt.Printf(format, i.Number, i.User.Login, i.Title)
		}
	}
}
