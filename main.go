package main

import "time"

func main() {
	maxURLsToCrawl := 50              // Set the maximum number of URLs to crawl
	crawlTimeout := 5 * time.Second   // Set the timeout for fetching URLs

	crawler := NewCrawler(maxURLsToCrawl, crawlTimeout)

	startURL := "https://books.toscrape.com" // Replace with the starting URL of your choice
	crawler.Frontier.AddURL(startURL)

	crawler.Crawl()
}

