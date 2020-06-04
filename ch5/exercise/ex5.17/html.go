package html

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if len(name) == 0 {
		return nil
	}
	if doc.Type == html.ElementNode {
		for _, tag := range name {
			if tag == doc.Data {
				nodes = append(nodes, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, name...)...)
	}
	return nodes
}
