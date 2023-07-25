package main

import (
	"fmt"
)

func main() {
	config, err:= NewConfig("config.json")
	if err != nil {
		fmt.Errorf("error reading the config file: %s", err)
	}
	// maxURLsToCrawl := 50              // Set the maximum number of URLs to crawl
	// crawlTimeout := 5 * time.Second   // Set the timeout for fetching URLs
	// maxReqeuestsPerSecond := 10

	fontier := NewURLFrontier()
	downloader := NewDownloader(config.OutputFolder, config.OutputName, config.OutputFileExtension)
	parser := NewParser(parseBook)
	limiter := NewRateLimiter(config.MaxRequestsPerSecond)

	crawler := NewCrawler(config.MaxURLsToCrawl, config.CrawlTimeoutSeconds, fontier, downloader, parser, limiter)

	// startURL := "https://books.toscrape.com/" // Replace with the starting URL of your choice
	crawler.Frontier.AddURL(config.StartURL)

	crawler.Crawl()
}


