package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

//base type for parseable struct
type Parseable interface {
	Parse(html, url string) (ParsedData, error)
}

// BaseParsedData contains common fields for parsed data
type ParsedData struct {
	Data []interface{}
}

// Parser struct
type Parser struct{
	Parser Parseable
}

func NewParser(item Parseable) *Parser {
	return &Parser{
		Parser: item,
	}
}

// Generic Parse function that accepts the parse function and HTML string to parse
func (p *Parser) Parse(html, url string) (ParsedData, error) {
	return p.Parser.Parse(html, url)
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

