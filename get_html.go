package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Pull all html from a website urling a GET request bytes are converted to a string
//
// Returns:
//  string - html body bytes
//  error - nil or error
func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("Network error: %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("HTTP error: %s\n", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("non-HTML response: %s\n", contentType)
	}

	htmlBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v\n", err)
	}

	htmlBody := string(htmlBodyBytes)

	return htmlBody, nil
}
