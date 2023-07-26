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
	UserAgents      []string
	UserAgentIndex  int
}

func NewCrawler(maxURLsToCrawl int, crawlTimeout time.Duration,  userAgents []string, fontier *URLFrontier, downloader *Downloader, parser *Parser, limiter *RateLimiter) *Crawler {
	return &Crawler{
		Frontier: fontier,
		Downloader: downloader,
		Parser: parser,
		Limiter: limiter,
		MaxURLsToCrawl:   maxURLsToCrawl,
		CrawlTimeout:     crawlTimeout,
		CrawledURLsCount: 0,
		UserAgents:      userAgents,
		UserAgentIndex:  0,
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

func (c *Crawler) processURL(url string) {
	if c.Frontier.HasURL(url) {
		fmt.Printf("Already visited url %s", url)
		return
	}

	// Set the user-agent for the request
	c.Downloader.SetUserAgent(c.UserAgents[c.UserAgentIndex])

	body, err := c.Downloader.Download(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %s\n", url, err)
		return
	}

	links := c.Parser.ExtractLinks(body, url) 
	fmt.Printf("Found %d links on %s\n", len(links), url)

	data, err := c.Parser.Parse(body)
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	err = c.Downloader.WriteDataToJSON(data, c.CrawledURLsCount)
	if err != nil {
		if err != nil {
			fmt.Println("Error downloading data:", err)
			return
		}
	}

	c.Frontier.RemoveURL(url)
	c.CrawledURLsCount++

	// Increment the user-agent index for the next request
	c.UserAgentIndex = (c.UserAgentIndex + 1) % len(c.UserAgents)

	for _, link := range links {
		c.Frontier.AddURL(link)
	}
}
