// Package github provides a Go API for the Github issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	APIURL    string = "https://api.github.com"
	IssuesURL        = "https://api.github.com/search/issues"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func getIssueURL(owner, repo, number string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, number)
}

func getIssuesURL(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
}

func setAuthorization(req *http.Request) error {
	// req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
	// My env cannot fix no token error,so ...
	os.Setenv("GITHUB_TOKEN", "<your_token_string>")
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return fmt.Errorf("GITHUB_TOKEN is not set")
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	return nil
}
