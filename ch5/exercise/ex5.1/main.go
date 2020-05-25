// Change the findlinks program to traverse the n.FirstChild linked list
// using recursive calls to visit instead of a loop.
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
		fmt.Fprintf(os.Stderr, "ex5.1: findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	/* for c := n.FirstChild; c != nil; c = c.NextSibling { */
	/*         links = visit(links, c) */
	/* } */
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
