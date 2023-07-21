package main

import "sync"

type URLFrontier struct {
	queue   []string
	visited map[string]bool
	mutex   sync.Mutex
}

func NewURLFrontier() *URLFrontier {
	return &URLFrontier{
		queue:   make([]string, 0),
		visited: make(map[string]bool),
	}
}

func (f *URLFrontier) AddURL(url string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if !f.visited[url] {
		f.queue = append(f.queue, url)
		f.visited[url] = true
	}
}

func (f *URLFrontier) GetNextURL() string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.queue) == 0 {
		return ""
	}

	nextURL := f.queue[0]
	f.queue = f.queue[1:]
	return nextURL
}

func (f *URLFrontier) IsEmpty() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return len(f.queue) == 0
}

func (f *URLFrontier) HasURL(url string) bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.visited[url]
}
