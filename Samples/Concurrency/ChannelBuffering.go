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
	//a channel of strings buffering up to 10 values.
	/**
	Because this channel is buffered,
	 we can send these values into the channel
	  without a corresponding concurrent receive.
	*/
	messages := make(chan string, 10)

	messages <- " Ping"
	messages <- " Pong"
	messages <- " Do you like it? "
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
