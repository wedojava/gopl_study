package main

import (
	"flag"
	"fmt"
	"log"

	LL "gopl.io/ch5/links"
)

var maxDepth int

type Link struct {
	url   string
	depth int
}

func crawl(l Link) []Link {
	fmt.Println(l.url)
	var links []Link
	if l.depth < maxDepth {
		depth := l.depth + 1
		list, err := LL.Extract(l.url)
		if err != nil {
			log.Print(err)
		}
		for _, url := range list {
			links = append(links, Link{url, depth})
		}
	}
	return links
}

func main() {
	flag.IntVar(&maxDepth, "depth", 3, "max crawl depth")
	flag.Parse()
	worklist := make(chan []Link)
	unseenLinks := make(chan Link)

	// Add command-line arguments to worklist.
	go func() {
		var links []Link
		for _, url := range flag.Args() {
			links = append(links, Link{url, 0})
		}
		worklist <- links
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
	close(unseenLinks)
}
