package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
    // Send unary to unblock channel
    cfg.concurrencyControl <- struct{}{}

	defer func() {
        // Mark done and block channel
        <-cfg.concurrencyControl
        cfg.wg.Done()
	}()

    if cfg.pagesLen() >= cfg.maxPages {
        return
    }

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// Exit if URL does not match expected baseURL
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
	}

    // Mutex to increment map
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

    // Get html content from the URL as a string
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

    // Traverse html body and find all <a> anchor tags with href tags
    // Returns all URLs from <a href="">
	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
	}

	for _, nextURL := range nextURLs {
        // Recursively add to waitgroup and crawl URL
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
