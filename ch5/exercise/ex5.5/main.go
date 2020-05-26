// Implement countWordsAndImages
// go run main.go https://xkcd.com
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ch05/ex05: %v\n", err)
			continue
		}
		fmt.Printf("words: %d\n", words)
		fmt.Printf("images: %d\n", images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return countWordsAndImages(n.NextSibling)
	}
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	cwords, cimages := countWordsAndImages(n.FirstChild)
	swords, simages := countWordsAndImages(n.NextSibling)
	words += cwords + swords
	images += cimages + simages
	return
}
