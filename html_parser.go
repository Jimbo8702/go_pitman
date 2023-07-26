package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// Define a generic interface for the struct to be parsed
type Parseable interface{}

// Define a generic parse function that takes a string and returns the parsed struct and error
type ParseFunc func(string) (Parseable, error)

// Parser struct
type Parser struct{
	Parser ParseFunc
}

func NewParser(parseFunc ParseFunc) *Parser {
	return &Parser{
		Parser: parseFunc,
	}
}

// Generic Parse function that accepts the parse function and HTML string to parse
func (p *Parser) Parse(html string) (Parseable, error) {
	return p.Parser(html)
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
