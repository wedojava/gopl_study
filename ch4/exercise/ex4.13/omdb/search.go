package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPoster(w io.Writer, terms []string) error {
	movie, err := search(terms)
	if err != nil {
		return err
	}
	if movie.Response {
		return fmt.Errorf("movie not found: %v", terms)
	}

	resp, err := http.Get(movie.Poster)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("get poster failed: %s", resp.Status)
	}

	_, err = io.Copy(w, resp.Body)
	resp.Body.Close()
	return err
}

func search(terms []string) (*Movie, error) {
	resp, err := http.Get(searchURL(terms))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get comic failed: %s", resp.Status)
	}

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
