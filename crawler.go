package main

import (
	"sync"
)

type Crawler struct {
	Visted map[string]bool
	mutex sync.Mutex
}

func NewCrawler() *Crawler {
	return &Crawler{
		Visted: make(map[string]bool),
	}
}