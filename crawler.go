package main

import (
	"fmt"
	"time"
)

type Crawler struct {
	Frontier *URLFrontier
	Downloader *Downloader
	Parser *Parser
	Limiter *RateLimiter
	MaxURLsToCrawl   int
	CrawlTimeout     time.Duration
	CrawledURLsCount int
}

func NewCrawler(maxURLsToCrawl int, crawlTimeout time.Duration, maxRequestsPerSecond int) *Crawler {
	return &Crawler{
		Frontier: NewURLFrontier(),
		Downloader: NewDownloader(),
		Parser: NewParser(),
		Limiter: NewRateLimiter(maxRequestsPerSecond),
		MaxURLsToCrawl:   maxURLsToCrawl,
		CrawlTimeout:     crawlTimeout,
		CrawledURLsCount: 0,
	}
}

func (c *Crawler) Crawl() {
	c.Limiter.init()

	go c.Limiter.run()

	for !c.Frontier.IsEmpty() {
		if c.CrawledURLsCount >= c.MaxURLsToCrawl {
			fmt.Println("Reached the maximum number of URLs to crawl.")
			break
		}

		// Wait for the rate limiter to allow processing the next URL
		c.Limiter.Wait()

		url := c.Frontier.GetNextURL()
		if url != "" {
			c.processURL(url)
		}
	}
}

// 
// fetches html and then extracts the links 
// marks the link as visited
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

	links := c.Parser.ExtractLinks(body, url) 
	fmt.Printf("Found %d links on %s\n", len(links), url)

	books, err := c.Parser.ParseBooks(body)
	if err != nil {
		fmt.Println("Error parsing books:", err)
		return
	}

	err = c.Downloader.WriteDataToJSON(books, c.CrawledURLsCount)
	if err != nil {
		if err != nil {
			fmt.Println("Error downloading data:", err)
			return
		}
	}

	c.Frontier.RemoveURL(url)
	c.CrawledURLsCount++

	for _, link := range links {
		c.Frontier.AddURL(link)
	}
}
