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
	maxURLsToCrawl := 50              // Set the maximum number of URLs to crawl
	crawlTimeout := 5 * time.Second   // Set the timeout for fetching URLs
	maxReqeuestsPerSecond := 10
	crawler := NewCrawler(maxURLsToCrawl, crawlTimeout,  maxReqeuestsPerSecond)

	startURL := "https://books.toscrape.com/" // Replace with the starting URL of your choice
	crawler.Frontier.AddURL(startURL)

	crawler.Crawl()
}

// func main() {
// 	// Create a rate limiter that allows 5 events per second
// 	limiter := NewRateLimiter(5)

// 	// Initialize the rate limiter and start the refill routine
// 	limiter.init()
// 	go limiter.run()

// 	// Simulate some events
// 	for i := 1; i <= 20; i++ {
// 		limiter.Wait()
// 		fmt.Printf("Event %d: %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
// 	}
// }


