package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateIssue(owner, repo string, fields map[string]string) error {
	buf := &bytes.Buffer{}
	// breakpoint here to study
	err := json.NewEncoder(buf).Encode(fields)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", getIssuesURL(owner, repo), buf)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Content-Type", "application/json")
	if err = setAuthorization(req); err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("post issue failed: %s", resp.Status)
	}
	return nil
}
