// Use the breadthFirst function to explore a different structure.
// For example, you could use the course dependencies from the topoSort example(a directed graph),
// the file system hierarchy on your computer(a tree), or a list of bus or subway routes
// downloaded from your city government's web site(an undirected graph).
// course preregs from ex5.10
package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func deps(course string) []string {
	fmt.Println(course)
	return prereqs[course]
}

func main() {
	var course string
	for course = range prereqs { // get random key
		break
	}
	breadthFirst(deps, []string{course})
}
