package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

// tokens is a counting semephore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	// tokens will block while len equal 20
	tokens <- struct{}{} //acquire a token
	list, err := links.Extract(url)
	<-tokens // relese the tokens so it's len < 20, process continue
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)

			}
		}
	}
}
