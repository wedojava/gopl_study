package main

import (
	"html/template"
)

func main() {
	var issueList = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} issues</h1>
	<table>
	    <tr>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	    </tr>
	    {{range .Items}}
	    <tr>
		<td><a href="{{.HTMLURL}}">{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	    </tr>
	    {{end}}
	</table>
	`))
}
