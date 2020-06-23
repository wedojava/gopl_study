// go run main.go div div h2
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	src := fetch("http://www.w3.org/TR/2006/REC-xml11-20060816")
	dec := xml.NewDecoder(bytes.NewBuffer(src))
	// dec := xml.NewDecoder(os.Stdin)
	var stack []string            // stack of element names
	var attrs []map[string]string // stack of element attributes
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			attr := make(map[string]string)
			for _, a := range tok.Attr {
				attr[a.Name.Local] = a.Value // key point line
			}
			attrs = append(attrs, attr)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			attrs = attrs[:len(attrs)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func fetch(url string) (rt []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	rt, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	// fmt.Printf("%s", b)
	return
}

// toStringSlice converts an element's name or attribute into a slice of representation for selecting an element.
func toStringSlice(stack []string, attrs []map[string]string) []string {
	r := []string{}
	for i := range stack {
		r = append(r, stack[i])
		for k, v := range attrs[i] {
			r = append(r, k+"="+v)
		}
	}
	return r
}
