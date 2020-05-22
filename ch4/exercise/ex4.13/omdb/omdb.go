package omdb

import (
	"fmt"
	"net/url"
	"strings"
)

type Movie struct {
	Poster   string
	Response string
}

func searchURL(terms []string) string {
	return fmt.Sprintf("http://www.omdbapi.com/?t=%s", url.QueryEscape(strings.Join(terms, " ")))
}
