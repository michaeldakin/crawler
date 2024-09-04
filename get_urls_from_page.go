package main

import (
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	page, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("get_urls_from_page.go - couldn't parse HTML: %v", err)
	}

	var urls []string
	var traverseNode func(*html.Node)
	traverseNode = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" { // grab all <a> anchors from page
			for _, anchor := range node.Attr {
				if anchor.Key == "href" { // grab all href tags from <a> anchors
					href, err := url.Parse(anchor.Val)
					if err != nil {
					    slog.Error("get_urls_from_page.go - couldn't parse href", "anchor", anchor.Val, "error", err)
						continue
					}

					resolvedURL := baseURL.ResolveReference(href) // absolute URI 
					urls = append(urls, resolvedURL.String())
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNode(child)
		}
	}
	traverseNode(page)

	return urls, nil
}
