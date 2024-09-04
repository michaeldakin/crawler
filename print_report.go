package main

import (
	"fmt"
	"sort"
)

type Link struct {
	URL  string
	Hits int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

    sorted := sortLinks(pages)
    for _, page := range sorted {
        url := page.URL
        hits := page.Hits
        fmt.Printf("Found %d internal links to %s\n", hits, url)
    }
}

func sortLinks(pages map[string]int) []Link {
    linksHitCount := []Link{}
    for url, hits := range pages {
        linksHitCount = append(linksHitCount, Link{url, hits})
    }

    sort.Slice(linksHitCount, func(i, j int) bool {
        if linksHitCount[i].Hits == linksHitCount[j].Hits {
            return linksHitCount[i].URL < linksHitCount[j].URL
        }
        return linksHitCount[i].Hits > linksHitCount[j].Hits
    })
    
    return linksHitCount
}
