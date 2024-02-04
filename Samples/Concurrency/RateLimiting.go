package main

import (
	"fmt"
	"time"
)

/*

Rate limiting is an important mechanism for
 controlling resource utilization and maintaining
  quality of service. Go elegantly supports rate limiting with goroutines, channels,
   and tickers
*/
func main() {
	/**
	Suppose we want to limit our handling of incoming requests.
	We’ll serve these requests off a channel of the same name.
	*/
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)
	/*
		This limiter channel will receive a value every 200 milliseconds.
		 This is the regulator in our rate limiting scheme
	*/
	limitter := time.Tick(1 * time.Second)
	for req := range requests {
		<-limitter
		fmt.Println("request", req, time.Now())
	}

}
