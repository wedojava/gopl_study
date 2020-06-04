package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
	"gopl.io/ch5/exercise/ex5.7_kdama/prettyhtml"
)

var depth int

// method: 4hel/gopl
// func ElementByID(n *html.Node, id string) *html.Node {
//         var retval *html.Node = nil
//         for _, a := range n.Attr {
//                 if a.Key == "id" && a.Val == id {
//                         retval = n
//                 }
//         }
//         for c := n.FirstChild; c != nil; c = c.NextSibling {
//                 if retval == nil {
//                         retval = getElementByID(c, id)
//                 }
//         }
//         return retval
// }

// method: torbiak/gopl: sth wrong!
// func ElementByID(n *html.Node, id string) *html.Node {
//         pre := func(n *html.Node) bool {
//                 if n.Type != html.ElementNode {
//                         return true
//                 }
//                 for _, a := range n.Attr {
//                         if a.Key == "id" && a.Val == id {
//                                 return false
//                         }
//                 }
//                 return true
//         }
//         return forEachNode(n, pre, pre)
// }
//
// func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
//         u := make([]*html.Node, 0) // unvisited
//         u = append(u, n)
//         for len(u) > 0 {
//                 n = u[0]
//                 u = u[1:]
//                 if pre != nil {
//                         if !pre(n) {
//                                 return n
//                         }
//                 }
//         }
//         for c := n.FirstChild; c != nil; c = c.NextSibling {
//                 u = append(u, c)
//         }
//         if post != nil {
//                 if !post(n) {
//                         return n
//                 }
//         }
//         return nil
// }

// method: kdama/gopl
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if !pre(n, id) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, id, pre, post)
		if node != nil {
			return node
		}
	}
	if post != nil {
		if !post(n, id) {
			return n
		}
	}
	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, findElement, nil)
}

// findElement find key and value to match the element, return false for matched and true for unmatched.
// it is logic reverse beacuse of forEachNode logic.
// the post func in forEachNode is useless for this metter.
func findElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
	}
	return true
}

func main() {
	// doc, err := html.Parse(os.Stdin)
	// if err != nil {
	//         fmt.Fprintln(os.Stderr, "usage: go run main.go < test.html")
	//         fmt.Fprintf(os.Stderr, "ex5.8: ElementByID: %v\n", err)
	//         os.Exit(1)
	// }
	// node := ElementByID(doc, "id1")
	// prettyhtml.WriteHTML(os.Stdout, node)

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: ex5.8 HTML_FILE ID")
	}
	filename := os.Args[1]
	id := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	prettyhtml.WriteHTML(os.Stdout, ElementByID(doc, id))
}
