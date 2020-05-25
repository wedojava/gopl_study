// Write a function to populate a mapping
// from element names -- p, div, span, and so on --
// to the number of elements with that name in an HTML document tree.
// go run gopl.io/ch1/fetch https://golang.org | go run main.go
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5.2: err: %v\n", err)
		os.Exit(1)
	}
	for k, v := range visit(make(map[string]int), doc) {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func visit(counts map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return counts
	}
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	visit(counts, n.FirstChild)
	visit(counts, n.NextSibling)
	return counts
}
