package main

import (
	"fmt"
	"time"
)

type Crawler struct {
	Frontier *URLFrontier
	Downloader *Downloader
	MaxURLsToCrawl   int
	CrawlTimeout     time.Duration
	CrawledURLsCount int
}

func NewCrawler(maxURLsToCrawl int, crawlTimeout time.Duration) *Crawler {
	return &Crawler{
		Frontier: NewURLFrontier(),
		Downloader: NewDownloader(),
		MaxURLsToCrawl:   maxURLsToCrawl,
		CrawlTimeout:     crawlTimeout,
		CrawledURLsCount: 0,
	}
}

func (c *Crawler) Crawl() {
	for !c.Frontier.IsEmpty() {
		if c.CrawledURLsCount >= c.MaxURLsToCrawl {
			fmt.Println("Reached the maximum number of URLs to crawl.")
			break
		}

		url := c.Frontier.GetNextURL()
		if url != "" {
			c.processURL(url)
		}
	}
}

// 
// fetches html and then extracts the links 
// marks the link as visited, and then dequeues it
//
func (c *Crawler) processURL(url string) {
	if c.Frontier.HasURL(url) {
		fmt.Printf("Already visited url %s", url)
		return
	}

	body, err := c.Downloader.Download(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %s\n", url, err)
		return
	}

	links := extractLinks(body, url) 
	fmt.Printf("Found %d links on %s\n", len(links), url)

	c.Frontier.RemoveURL(url)
	c.CrawledURLsCount++

	for _, link := range links {
		c.Frontier.AddURL(link)
	}
}
