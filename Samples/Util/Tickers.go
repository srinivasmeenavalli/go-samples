package main

import (
	"fmt"
	"time"
)

/*

Timers are for when you want to do something
once in the future - tickers are for when you
 want to do something repeatedly at regular intervals
*/
func main() {
	/*
		Tickers use a similar mechanism to timers:
		a channel that is sent values. Here we’ll
		 use the select builtin on the channel to await the values as they arrive
		  every second
	*/
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Ticker at T", t)
			}
		}

	}()
	/*
		Tickers can be stopped like timers.
		 Once a ticker is stopped it won’t receive
		 any more values on its channel. We’ll stop ours after 10 seconds
	*/
	time.Sleep(10 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticket Stopped")
}
