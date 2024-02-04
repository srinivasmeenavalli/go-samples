package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		   allow short bursts of requests in our rate limiting scheme
		    while preserving the overall rate limit. We can accomplish
			this by buffering our limiter channel.
			This burstyLimiter channel will allow bursts of up to 3 events.
	*/
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	/*Every 2 seconds we’ll try to
	add a new value to burstyLimiter, up to its limit of 3
	*/

	go func() {
		for t := range time.Tick(2 * time.Second) {
			burstyLimiter <- t
		}

	}()
	/*
		Now simulate 15 more incoming requests.
		 The first 3 of these will benefit from the burst capability of burstyLimiter
	*/
	burstyRequests := make(chan int, 15)

	for i := 1; i <= 15; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	/*
		serve the first 3 immediately because of the burstable rate
		limiting, then serve the remaining 12 with ~2 seconds delays each
	*/
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())

	}
}
