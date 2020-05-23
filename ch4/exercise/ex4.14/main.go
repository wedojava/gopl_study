// Create a web server that queries GitHub once
// and then allows navigation of the list of
// bug reports, milestones, and users.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"gopl.io/ch4/exercise/ex4.14/github"
)

var navigation = `
<p>
<a href='/'>Issues</a>
<a href='/milestones'>Milestones</a>
<a href='/users'>Users</a>
`

var issueTable = `
<h1>{{len .}} issue{{if ne (len .) 1}}s{{end}}</h1>
<table>
    <tr style = "text-align: left">
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
    </tr>
    {{range .}}
    <tr>
	<td><a href="{{.HTMLURL}}">{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href="{{.User.HTMLURL}}">{{.User.Login}}</a></td>
	<td><a href="{{.HTMLURL}}">{{.Title}}</a></td>
    </tr>
    {{end}}
</table>
`

var milestoneTable = `
<h1>{{len .}} issue{{if ne (len .) 1}}s{{end}}</h1>
<table>
    <tr style = "text-align: left">
	<th>Title</th>
	<th>State</th>
    </tr>
    {{range .}}
    <tr>
	<td><a href = "{{.HTMLURL}}">{{.Title}}</a></td>
	<td>{{.State}}</td>
    </tr>
    {{end}}
</table>
`

var userTable = `
<h1>{{len .}} user{{if ne (len .) 1}}s{{end}}</h1>
<table>
    <tr style = "text-align: left">
	<th>Avatar</th>
	<th>Username</th>
    </tr>
    {{range .}}
    <tr>
	<td><a href="{{.HTMLURL}}"><img src="{{.AvatarURL}}" width="32" heigh="32"></td>
	<td><a href="{{.HTMLURL}}">{{.Login}}</a></td>
    </tr>
    {{end}}
</table>
`

var issuesTemplate = template.Must(template.New("issues").Parse(navigation + issueTable))
var milestonesTemplate = template.Must(template.New("milestones").Parse(navigation + issueTable))
var usersTemplate = template.Must(template.New("users").Parse(navigation + userTable))
var issues []github.Issue
var milestones []github.Milestone
var users []github.User

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: ex14 OWNER REPO")
		os.Exit(1)
	}
	owner, repo := os.Args[1], os.Args[2]
	if err := generateCache(owner, repo); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handleIssues)
	http.HandleFunc("/milestones", handleMilestones)
	http.HandleFunc("/users", handleUsers)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generateCache(owner, repo string) error {
	issues, err := github.GetIssues(owner, repo)
	if err != nil {
		return err
	}
	for _, issue := range issues {
		if issue.Milestone != nil {
			milestones = appendMilestoneAsSet(milestones, issue.Milestone)
		}
		for _, assignee := range issue.Assignees {
			users = appendUserAsSet(users, assignee)
		}
		users = appendUserAsSet(users, issue.User)
	}
	return nil
}

func handleIssues(w http.ResponseWriter, r *http.Request) {
	issuesTemplate.Execute(w, issues)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	usersTemplate.Execute(w, users)
}

func handleMilestones(w http.ResponseWriter, r *http.Request) {
	milestonesTemplate.Execute(w, milestones)
}

func appendMilestoneAsSet(set []github.Milestone, milestone *github.Milestone) []github.Milestone {
	if !includesMilestone(set, milestone) {
		return append(set, *milestone)
	}
	return set
}

func includesMilestone(array []github.Milestone, milestone *github.Milestone) bool {
	for _, v := range array {
		if v.Equals(milestone) {
			return true
		}
	}
	return false
}

func appendUserAsSet(set []github.User, user *github.User) []github.User {
	if !includesUser(set, user) {
		return append(set, *user)
	}
	return set
}

func includesUser(array []github.User, user *github.User) bool {
	for _, v := range array {
		if v.Equals(user) {
			return true
		}
	}
	return false
}
