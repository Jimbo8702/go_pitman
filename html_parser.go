package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

func (p *Parser) ParseBooks(html string) ([]Book, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var books []Book

	doc.Find("article.product_pod").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3 a").Text()
		price := s.Find("p.price_color").Text()
		availability := s.Find("p.instock.availability").Text()

		availability = strings.TrimSpace(availability)

		book := Book{
			Title:       title,
			Price:       price,
			Availability: availability,
		}

		books = append(books, book)
	})

	return books, nil
}