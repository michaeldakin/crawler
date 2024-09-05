## About the project

This repo is used to track the boot.dev "Build a Web Crawler" course.

Usage: ./crawler --url "example.com" -mc 2 -mp 25
- URL to scan | -url string
- Max goroutines - | -mc int
- Max pages to scan - | -mp int


Features:
* Uses recursion to crawl a website
* Uses goroutines to crawl individual links, limited by maxConcurrency (-mc) flag
* I suck at readme's
