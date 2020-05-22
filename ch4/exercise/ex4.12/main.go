package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopl.io/ch4/exercise/ex4.12/xkcd"
)

const maxComicID = 2309

var fetchFlag = flag.Bool("fetch", false, "fetch all comics")

func main() {
	flag.Parse()

	if *fetchFlag {
		fetch()
	} else {
		if len(flag.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "ch04/exercise/ex4.12 must have at least 1 query")
			os.Exit(1)
		}
		search(flag.Args())
	}
}

func fetch() {
	comicIndex := xkcd.NewComicIndex()
	for comicID := 1; comicID <= maxComicID; comicID++ {
		comic, err := xkcd.GetComic(comicID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ch04/exercise/ex4.12: get %d: %v", comicID, err)
		} else {
			comicIndex.ComicSlice = append(comicIndex.ComicSlice, comic)
		}
	}
	result, err := json.MarshalIndent(comicIndex, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/exercise/ex4.12: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", result)
	}
}

func search(terms []string) {
	index, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/exercise/ex4.12: %v\n", err)
		os.Exit(1)
	}

	ci := xkcd.NewComicIndex()
	err = json.Unmarshal(index, &ci)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/exercise/ex4.12: %v\n", err)
		os.Exit(1)
	}
	searchResult := xkcd.SearchComic(ci, terms)

	fmt.Printf("%d comics:\n", len(searchResult))
	for _, c := range searchResult {
		printComic(c)
	}
}

func printComic(c *xkcd.Comic) {
	fmt.Printf("\n-- Comic %d --\n", c.Num)
	fmt.Printf("\nImage URL:\n%s\n", c.Img)
	fmt.Printf("\nTranscript:\n%s\n", c.Transcript)
}
