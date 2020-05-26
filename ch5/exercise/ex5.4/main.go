// Extend the visit function so that
// it extracts other kinds of links from the document,
// such as images, scripts, and style sheets.
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
		fmt.Fprintf(os.Stderr, "ch05/ex04: %v\n", err)
		os.Exit(1)
	}
	links := visit(nil, doc)
	for _, link := range links {
		fmt.Println(link)
	}

}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode &&
		(n.Data == "a" || n.Data == "img" || n.Data == "script" || n.Data == "video" || n.Data == "link") {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" || a.Key == "poster" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

// var linkAttrs = map[string][]string{
//         "a":      []string{"href"},
//         "link":   []string{"href"},
//         "img":    []string{"src"},
//         "script": []string{"src"},
//         "iframe": []string{"src"},
//         "form":   []string{"action"},
//         "html":   []string{"manifest"},
//         "video":  []string{"src", "poster"},
// }
// func visit(links []string, n *html.Node) []string {
//         if n.Type == html.ElementNode {
//                 for key, attrs := range linkAttrs {
//                         if n.Data == key {
//                                 for _, attr := range attrs {
//                                         for _, a := range n.Attr {
//                                                 if a.Key == attr {
//                                                         links = append(links, a.Val)
//                                                 }
//                                         }
//                                 }
//                         }
//                 }
//         }
//         for c := n.FirstChild; c != nil; c = c.NextSibling {
//                 links = visit(links, c)
//         }
//         return links
// }
