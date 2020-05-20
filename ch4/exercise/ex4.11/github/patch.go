package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func patchIssue(owner, repo, number string, fields map[string]string) error {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf) // so how's it going?
	if err := encoder.Encode(fields); err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", getIssueURL(owner, repo, number), buf)
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

	resp.Body.Close() // without defer?

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("pathc issue failed: %s", resp.Status)
	}
	return nil
}

func UpdateIssue(owner, repo, number string, fields map[string]string) error {
	return patchIssue(owner, repo, number, fields)
}

func ReopenIssue(owner, repo, number string) error {
	fields := map[string]string{
		"state": "open",
	}
	return patchIssue(owner, repo, number, fields)
}

func CloseIssue(owner, repo, number string) error {
	fields := map[string]string{
		"state": "closed",
	}
	return patchIssue(owner, repo, number, fields)
}
