// Package github provides a Go API for the Github issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"fmt"
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
	UpdatedAt time.Time
	Body      string
	Assignees []*User
	Milestone *Milestone
}

type User struct {
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
	ID        int
	Login     string
}

type Milestone struct {
	Description string
	HTMLURL     string `json:"html_url"`
	ID          int
	State       string
	Title       string
}

func getIssuesURL(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?state=all", owner, repo)
}

func (u1 *User) Equals(u2 *User) bool {
	return u1.ID == u2.ID
}

func (m1 *Milestone) Equals(m2 *Milestone) bool {
	return m1.ID == m2.ID
}
