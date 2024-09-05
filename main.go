package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

var (
	flagURL            string
	flagMaxConcurrency int
	flagMaxPages       int
)

func init() {
	flag.StringVar(&flagURL, "url", "", "URL to crawl")
	flag.IntVar(&flagMaxConcurrency, "mc", 1, "Amount of goroutines to use, use a lower number to prevent rate limting")
	flag.IntVar(&flagMaxPages, "mp", 10, "Max pages to crawl")
}

func main() {
	flag.Parse()
	fmt.Printf("found %d flags\n", flag.NFlag())


	if flag.NFlag() < 3 {
		fmt.Println("not enough args")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	if flag.NFlag() > 3 {
		fmt.Println("too many args provided")
		os.Exit(1)
	}

	if flagURL == "" {
		fmt.Println("No URL provided")
		os.Exit(1)
	}

	if flagMaxConcurrency >= runtime.NumCPU() {
        fmt.Println("numCPUs:", runtime.NumCPU())
		fmt.Println("Too many goroutines for this CPU - lower maxConcurrency")
		os.Exit(1)
	}

	if flagMaxConcurrency < 1 {
		fmt.Println("Not enough goroutines - minimum maxConcurrency is 1")
		os.Exit(1)
	}

	if flagMaxPages < 1 {
		fmt.Println("Not enough maxPages - minimum is 1 (but you probably want at more)")
		os.Exit(1)
	}

	if flagMaxPages > 100 {
		fmt.Println("That is a lot of pages to crawl - this may cause Rate Limiting or other action")
	}

	cfg, err := configure(flagURL, flagMaxConcurrency, flagMaxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s...\n\n", flagURL)
	startTime := time.Now()

	cfg.wg.Add(1)
	cfg.crawlPage(flagURL)
	cfg.wg.Wait()

	printReport(cfg.pages, flagURL)

	fmt.Printf("crawling took %v seconds\n", time.Since(startTime))
}
