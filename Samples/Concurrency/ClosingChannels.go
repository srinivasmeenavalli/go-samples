package main

/*

Closing a channel indicates that no more values will be sent on it
*/
import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs

			if more {
				fmt.Println("received job No:", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}

		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		time.Sleep(time.Second)
	}
	fmt.Println("sent all jobs")
	//jobs <- 4
	//jobs <- 5
	/*
		use a jobs channel to communicate work to be done
		 from the main() goroutine to a worker goroutine.
		 When we have no more jobs for the worker we’ll close the jobs channel
	*/
	close(jobs)

	<-done
}
