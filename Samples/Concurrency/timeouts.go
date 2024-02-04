package main

import (
	"fmt"
	"time"
)

/*
Timeouts are important for programs that connect
to external resources or that otherwise need to bound execution time.
*/
func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		/**
		channel is buffered,
		so the send in the goroutine is nonblocking
		*/
		c1 <- "result 1"
	}()
	/*
	 Since select proceeds with the first receive that’s ready,
	 we’ll take the timeout case if the operation
	 takes more than the allowed 1s
	*/
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout Happend after 1 second")
	}
	//Results returned before Timeout
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Response with in 2 seconds"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout Happend after 3 second")
	}
}
