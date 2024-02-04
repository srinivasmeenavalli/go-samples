package main

/*
We can also use this syntax to iterate over values received from a channel
*/
import "fmt"

func main() {
	queue := make(chan string, 5)
	queue <- "Job One"
	queue <- "Job Two"
	queue <- "Job Three"
	close(queue)
	/*
		Range iterates over each element
		as it’s received from queue. Because we closed the channel above,
		 the iteration terminates after receiving the 3 elements
	*/
	for elem := range queue {
		fmt.Println(elem)
	}

}
