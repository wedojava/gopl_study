package xkcd

import (
	"fmt"
)

type Comic struct {
	Num        int
	Link       string
	Year       string
	Month      string
	Day        string
	News       string
	SafeTitle  string
	Transcript string
	Alt        string
	Img        string
	Title      string
}

// const usage = `
// xkcd get N
// xkcd index OUTPUT_FILE
// xkcd search INDEX_FILE QUERY
// `

type ComicIndex struct {
	ComicSlice []*Comic
}

func getComicURL(comicID int) string {
	return fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicID)
}

// func UsageDie() {
//         fmt.Println(usage)
//         os.Exit(1)
// }

func NewComicIndex() ComicIndex {
	return ComicIndex{[]*Comic{}}
}
