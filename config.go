package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
    concurrencyControl chan struct{}
    maxPages           int
	pages              map[string]int
    baseURL            *url.URL
	mu                 *sync.Mutex
	wg                 *sync.WaitGroup
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse baseURL: %v", err)
	}

	return &config{
        concurrencyControl: make(chan struct{}, maxConcurrency),
        maxPages:           maxPages,
		pages:              make(map[string]int),
        baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
	}, nil
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()

	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return false
	}

	// add first visit to map
	cfg.pages[normalizedURL] = 1

    cfg.mu.Unlock()
	return true
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}
