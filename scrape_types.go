package main

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Book struct {
	URL 		 string
	Title        string
	Price        string
	Availability string
	CreatedAt 	 time.Time
}

func (b Book) Parse(html, url string) (ParsedData, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ParsedData{}, err
	}

	var books []Book

	doc.Find("article.product_pod").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3 a").Text()
		price := s.Find("p.price_color").Text()
		availability := s.Find("p.instock.availability").Text()

		availability = strings.TrimSpace(availability)

		book := Book{
			URL :		 url,
			Title:       title,
			Price:       price,
			Availability: availability,
			CreatedAt: time.Now().UTC(),
		}

		books = append(books, book)
	})

	// Perform a type conversion to []interface{} for the books slice
	var dataSlice []interface{}
	for _, b := range books {
		dataSlice = append(dataSlice, b)
	}

	parsedData := ParsedData{
		Data: dataSlice,
	}

	return parsedData, nil
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

func (j *Jet) Parse(html, url string) (ParsedData, error) {
	return ParsedData{}, nil
}