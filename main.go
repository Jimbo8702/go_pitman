package main

import "fmt"

func main() {
	config, err:= NewConfig("config.json")
	if err != nil {
		fmt.Errorf("error reading the config file: %s", err)
	}

	fontier := NewURLFrontier(config.StartURL)
	downloader := NewDownloader(config.OutputFolder, config.OutputName, config.OutputFileExtension)
	parser := NewParser(parseBook)
	limiter := NewRateLimiter(config.MaxRequestsPerSecond)

	crawler := NewCrawler(config.MaxURLsToCrawl, config.CrawlTimeoutSeconds, config.UserAgents, fontier, downloader, parser, limiter)

	crawler.Crawl()
}


