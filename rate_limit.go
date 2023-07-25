package main

import (
	"time"
)

type RateLimiter struct {
	rate   int           // Number of events allowed per second
	bucket chan struct{} // Channel to hold the tokens (1 token = 1 event)
}

func NewRateLimiter(rate int) *RateLimiter {
	return &RateLimiter{
		rate:   rate,
		bucket: make(chan struct{}, rate),
	}
}

func (rl *RateLimiter) init() {
	// Fill the bucket with tokens initially
	for i := 0; i < rl.rate; i++ {
		rl.bucket <- struct{}{}
	}
}

func (rl *RateLimiter) run() {
	// Refill the bucket periodically
	for range time.Tick(time.Second) {
		// Refill the bucket with tokens
		for i := 0; i < rl.rate; i++ {
			select {
			case rl.bucket <- struct{}{}:
			default:
				// If the bucket is full, break out of the loop
				break
			}
		}
	}
}

func (rl *RateLimiter) Wait() {
	// Acquire a token from the bucket (blocking if the bucket is empty)
	<-rl.bucket
}
