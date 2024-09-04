package main

import (
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

var count atomic.Uint32

func main() {
	if len(os.Args) < 4 {
		fmt.Println("not enough args")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		return
	}

	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]
	maxConcurrencyStr := os.Args[2]
	maxPagesStr := os.Args[3]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyStr)
	if err != nil {
		fmt.Printf("Error - maxConcurrency %v\n", err)
		return
	}
	maxPages, err := strconv.Atoi(maxPagesStr)
	if err != nil {
		fmt.Printf("Error - maxPages %v\n", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxPages, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	startTime := time.Now()

	cfg.wg.Add(1)
	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

    printReport(cfg.pages, rawBaseURL)

	fmt.Printf("crawling took %v seconds\n", time.Since(startTime))
}
