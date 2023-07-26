package main

import "fmt"

// const corePrompt = "Take the this html: {html} and this golang struct: {struct}, and build me a function similar to : {function}, using this new html and struct."

func main() {
	config, err:= NewConfig("config.json")
	if err != nil {
		fmt.Println("error reading the config file: %s", err)
	}

	downloader := NewDownloader(config.OutputFolder, config.OutputName, config.OutputFileExtension)
	builder := NewStructBuilder(config.SchemaName, config.SchemaModelFile, config.ModelsFolderName)
	
	structString := builder.GenerateStructDefinition(config.Schema)
	fmt.Printf("Your new struct: %s", structString)

	//downloader fetches inital html ? 
	body, err := downloader.Download(config.StartURL)
	if err != nil {
		fmt.Printf("Error fetching START_URL %s: %s\n", config.StartURL, err)
	}

	//trim html 
	html := trimHTML(body)
	fmt.Printf("Your html string: %s", html)

	//send to chatgpt to fetch tags for the following struct
		//build prompt 
		//send to completion
		//trim string to just function definitation
		//build function definitation
		//build function
		//write golang file with struct and parse function
		//take function pointer and return it
		//initate parser with function
		//begin scraping
		

	err = builder.WriteGoFile(structString)
	if err != nil {
		fmt.Println("Error building Go file:", err)
	}
}

// fontier := NewURLFrontier(config.StartURL)
// downloader := NewDownloader(config.OutputFolder, config.OutputName, config.OutputFileExtension)
// parser := NewParser(parseBook)
// 	limiter := NewRateLimiter(config.MaxRequestsPerSecond)

// 	crawler := NewCrawler(config.MaxURLsToCrawl, config.CrawlTimeoutSeconds, fontier, downloader, parser, limiter)

// 	crawler.Crawl()
