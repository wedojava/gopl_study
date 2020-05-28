package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var depth int

func getElementByID(n *html.Node, id string) *html.Node {
	var retval *html.Node = nil
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			retval = n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if retval == nil {
			retval = getElementByID(c, id)
		}
	}
	return retval
}

func ElementByID(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
		return true
	}
	return forEachNode(n, pre, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0) // unvisited
	u = append(u, n)
	for len(u) > 0 {
		n = u[0]
		u = u[1:]
		if pre != nil {
			if !pre(n) {
				return n
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		u = append(u, c)
	}
	if post != nil {
		if !post(n) {
			return n
		}
	}
	return nil
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "usage: go run main.go < test.html")
		fmt.Fprintf(os.Stderr, "ex5.8: ElementByID: %v\n", err)
		os.Exit(1)
	}
	node := getElementByID(doc, "id1")
	fmt.Println(*node)
	// if len(os.Args) != 3 {
	//         fmt.Fprintf(os.Stderr, "usage: ex5.8 HTML_FILE ID")
	// }
	// filename := os.Args[1]
	// id := os.Args[2]
	// file, err := os.Open(filename)
	// if err != nil {
	//         log.Fatal(err)
	// }
	// doc, err := html.Parse(file)
	// if err != nil {
	//         log.Fatal(err)
	// }
	// fmt.Println(ElementByID(doc, id))
}
