package prettyhtml

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func WriteHTML(w io.Writer, n *html.Node) {
	forEachNode(w, n, startElement, endElement)
}

func forEachNode(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {
	if pre != nil {
		pre(w, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}
	if post != nil {
		post(w, n)
	}
}

var depth int

func startElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		if len(n.Attr) > 0 {
			for _, attr := range n.Attr {
				fmt.Fprintf(w, " %s=%q", attr.Key, attr.Val)
			}
		}
		if isSelfClosableTag(n.Data) && n.FirstChild == nil {
			fmt.Fprintf(w, "/>\n")
		} else {
			fmt.Fprintf(w, ">\n")
		}
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Fprintf(w, "%*s<!--%s-->\n", depth*2, "", n.Data)
	} else if n.Type == html.TextNode {
		re := regexp.MustCompile(`^[ \t ]*$`)
		for _, str := range strings.Split(n.Data, "\n") {
			if !re.MatchString(str) {
				fmt.Fprintf(w, "%*s%s\n", depth*2, "", strings.TrimSpace(str))
			}
		}
	}
}

func endElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if !isSelfClosableTag(n.Data) || n.FirstChild != nil {
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}

}

func isSelfClosableTag(tagName string) bool {
	return tagName == "area" ||
		tagName == "base" ||
		tagName == "br" ||
		tagName == "col" ||
		tagName == "command" ||
		tagName == "embed" ||
		tagName == "hr" ||
		tagName == "img" ||
		tagName == "input" ||
		tagName == "keygen" ||
		tagName == "link" ||
		tagName == "meta" ||
		tagName == "param" ||
		tagName == "source" ||
		tagName == "track" ||
		tagName == "wbr"
}
