package main

/*

We can use channels to synchronize execution across goroutines
*/

import (
	"fmt"
	"time"
)

/*
This is the function we’ll run in a goroutine.
The done channel will be used to notify another
goroutine that this function’s work is done
*/
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("Done")
	//Send a value to notify that we’re done
	done <- true

}
func main() {
	done := make(chan bool, 1)

	go worker(done)
	/**
	Block until we receive a notification from the worker on the channel
	If you removed the <- done
	line from this program, the program would
	 exit before the worker even started

	*/
	<-done

}
