package xkcd

import (
	"strings"
)

func SearchComic(from ComicIndex, terms []string) []*Comic {
	result := []*Comic{}

	for _, comic := range from.ComicSlice {
		if hit(comic, terms) {
			result = append(result, comic)
		}
	}

	return result
}

func hit(comic *Comic, terms []string) bool {
	for _, term := range terms {
		switch {
		case strings.Contains(comic.Alt, term):
			continue
		case strings.Contains(comic.Day, term):
			continue
		case strings.Contains(comic.Img, term):
			continue
		case strings.Contains(comic.Link, term):
			continue
		case strings.Contains(comic.Month, term):
			continue
		case strings.Contains(comic.News, term):
			continue
		case strings.Contains(comic.SafeTitle, term):
			continue
		case strings.Contains(comic.Title, term):
			continue
		case strings.Contains(comic.Transcript, term):
			continue
		case strings.Contains(comic.Year, term):
			continue
		default:
			return false
		}
	}
	return true
}
