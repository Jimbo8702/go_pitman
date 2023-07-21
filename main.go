package main

import "time"

// books := []Book{
// 	{Title: "fancy title", Price:"90", Availability:"yes"},
// 	{Title: "fancy title", Price:"90", Availability:"yes"},
// 	{Title: "fancy title", Price:"90", Availability:"yes"},

// }

// downloader := NewDownloader()

// downloader.WriteDataToJSON(books, 1)

func main() {
	maxURLsToCrawl := 10              // Set the maximum number of URLs to crawl
	crawlTimeout := 5 * time.Second   // Set the timeout for fetching URLs

	crawler := NewCrawler(maxURLsToCrawl, crawlTimeout)

	startURL := "https://books.toscrape.com" // Replace with the starting URL of your choice
	crawler.Frontier.AddURL(startURL)

	crawler.Crawl()
}

