package main

import (
	"fmt"
	"time"
)

/*
hese workers will receive work on the jobs channel
 and send the corresponding results on results.
  We’ll sleep a second per job to simulate an expensive task
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker ", id, "Started Job", j)
		time.Sleep(2 * time.Second)
		fmt.Println("Worker ", id, "Finished Job", j)
		results <- j * 2
	}
}
func main() {

	const numjobs = 5
	/*In order to use our pool of workers
	we need to send them work and collect their results
	*/
	jobs := make(chan int, numjobs)
	results := make(chan int, numjobs)
	/*
		This starts up 3 workers, initially blocked
		because there are no jobs ye
	*/
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	/*
		we send 5 jobs and then close that channel to indicate that’s all the work we have
	*/
	for j := 1; j <= numjobs; j++ {
		jobs <- j
	}
	close(jobs)
	/*
		Finally we collect all the results of the work.
		This also ensures that the worker goroutines have finished
	*/
	for a := 1; a <= numjobs; a++ {
		<-results
	}

}
