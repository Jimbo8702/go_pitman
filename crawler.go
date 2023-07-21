package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Crawler struct {
	Visited map[string]bool
	mutex sync.Mutex
}

func NewCrawler() *Crawler {
	return &Crawler{
		Visited: make(map[string]bool),
		mutex: sync.Mutex{},
	}
}

func (c *Crawler) fetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch %s: %s", url, resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func (c *Crawler) extractLinks(body string) []string {
	var links []string
	tokenizer := html.NewTokenizer(strings.NewReader(body))

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}

	return links
}

func (c *Crawler) Crawl(url string) {
	c.mutex.Lock()
	if c.Visited[url] {
		c.mutex.Unlock()
		return
	}
	c.Visited[url] = true
	c.mutex.Unlock()

	body, err := c.fetchHTML(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %s\n", url, err)
		return
	}

	links := c.extractLinks(body)
	fmt.Printf("Found %d links on %s\n", len(links), url)

	var wg sync.WaitGroup
	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			c.Crawl(link)
		}(link)
	}
	wg.Wait()
}



