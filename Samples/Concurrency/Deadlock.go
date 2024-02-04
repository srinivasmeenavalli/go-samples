package main

import "fmt"

/**

By default channels are unbuffered,
 meaning that they will only accept sends (chan <-)
 if there is a corresponding receive (<- chan)
 ready to receive the sent value.
 Buffered channels
 accept a limited number of values
  without a corresponding receiver for those values.
*/
func main() {
	//a channel of strings buffering up to 2 values.
	/**
	Because this channel is buffered,
	 we can send 2 values into the channel
	  without a corresponding concurrent receiver
	*/
	messages := make(chan string, 2)

	messages <- " Ping"
	messages <- " Pong"
	//fatal error: all goroutines are asleep - deadlock!
	//Deadlock will occuer because of buffer size
	messages <- " Do you like it? "

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
