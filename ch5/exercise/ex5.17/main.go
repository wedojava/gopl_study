package main

import "golang.org/x/net/html"

func main() {

}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var es []*html.Node

	return es
}

func forEachNode(n *html.Node, tag string, pre, post func(n *html.Node, tag string) bool) *html.Node {
	if pre != nil {
		if !pre(n, tag) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, tag, pre, post)
		if node != nil {
			return node
		}
	}
	if post != nil {
		if !post(n, tag) {
			return n
		}
	}
	return nil
}
