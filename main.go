package main

import "fmt"

func main() {
	config, err:= NewConfig("config.json")
	if err != nil {
		fmt.Errorf("error reading the config file: %s", err)
	}

	// builder := NewStructBuilder(config.SchemaName, config.SchemaModelFile, config.ModelsFolderName)
	
	// err = builder.GenerateStruct(config.Schema)
	// if err != nil {
	// 	fmt.Println("Error building Go file:", err)
	// 	return
	// }

	fontier := NewURLFrontier(config.StartURL)
	downloader := NewDownloader(config.OutputFolder, config.OutputName, config.OutputFileExtension)
	parser := NewParser(parseBook)
	limiter := NewRateLimiter(config.MaxRequestsPerSecond)

	crawler := NewCrawler(config.MaxURLsToCrawl, config.CrawlTimeoutSeconds, fontier, downloader, parser, limiter)

	crawler.Crawl()
}

//if no ai 
//we take a struct with tags 
//take those tags and build a parse function to take the given takes and struct and extract the data from the html
// pass the parse function to the parser and build the crawler

//if ai
//we need to 
//take the start url 
//call one call
//take the html string and send to chatgpt

//option - one
//tell chatgpt to build the function to parse the given struct
//then pass that function to parser and build the crawler
//return the crawler 

//option - two
//tell chatgpt to take the html and the given struct, and match the tags like [structField]:[data-tag]
//then pass that function to parser and build the crawler
//return the crawler 

//for the given struct
//"schema" field in config
//take schema and build a struct 
//build a tag struct and a data struct
