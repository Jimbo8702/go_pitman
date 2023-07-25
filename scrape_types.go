package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Book struct {
	Title       string
	Price       string
	Availability string
}

func parseBook(html string) (Parseable, error) {
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

type Jet struct {
	Manufacturer string
	Model string
	Category string
	MaxPassengerCapacity int
	MaxRangeMiles int 
	MaxSpeedMPH int
	CabinWidth string
	CabinLength string
	CabinHeight string
	BaggageCapacity string
	Lavatory string
	PriceUSD string
	PriceNote string
}