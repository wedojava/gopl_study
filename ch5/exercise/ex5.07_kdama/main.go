// go run main.go https://gopl.io
package main

import (
	"net/http"
	"os"

	"golang.org/x/net/html"
	"gopl.io/ch5/exercise/ex5.07_kdama/prettyhtml"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	prettyhtml.WriteHTML(os.Stdout, doc)
	return nil
}
