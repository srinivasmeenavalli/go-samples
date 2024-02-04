package main

import "fmt"

/**
Channels are the pipes that connect concurrent goroutines.
*/

func main() {
	/**
	Create a new channel with make(chan val-type). Channels are typed by the values they convey
	*/

	messages := make(chan string)
	go func() {
		//Send a value into a channel using the channel <- syntax.cls

		messages <- "ping"
		messages <- "pong"
	}()

	msg1 := <-messages
	msg2 := <-messages
	/**
	By default sends and receives block until both the sender and receiver are ready
	*/
	fmt.Println("Received message: ", msg1)
	fmt.Println("Received message: ", msg2)
}
