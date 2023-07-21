package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ExtractLinks(body string, baseURL string) []string {
	var links []string
	tokenizer := html.NewTokenizer(strings.NewReader(body))

	base, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing baseURL:", err)
		return links
	}

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
						linkURL, err := base.Parse(attr.Val)
						if err == nil {
							links = append(links, linkURL.String())
						}
					}
				}
			}
		}
	}

	return links
}