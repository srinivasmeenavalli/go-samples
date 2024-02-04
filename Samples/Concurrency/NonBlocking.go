package main

import "fmt"

/*
Basic sends and receives on channels are blocking. However,
we can use select with a default clause to implement non-blocking sends, receives
*/
func main() {
	messages := make(chan string, 2)
	signals := make(chan bool, 2)
	msg := "Hello"
	signals <- true
	select {
	case messages <- msg:
		fmt.Println("Sent Message", msg)
	default:
		fmt.Println("Non blocking: No message sent")
	}

	select {
	case msg := <-messages:
		//value is available on messages
		fmt.Println("received message:", msg)
	default:
		// If not it will immediately take the default case
		fmt.Println("Non blocking: No message received")
	}
	select {
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
