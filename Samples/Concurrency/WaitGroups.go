package main

/*
To wait for multiple goroutines to finish, we can use a wait group.
*/
import (
	"fmt"
	"sync"
	"time"
)

/*
function we’ll run in every goroutine.
*/
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	//Sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {
	/*
		This WaitGroup is used to wait for all the goroutines launched here to finish.
		Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	*/
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		i := i
		/*
			Wrap the worker call in a closure that makes sure to tell
			 the WaitGroup that this worker is done.
			 This way the worker itself does not have to be
			 aware of the concurrency primitives involved in its execution
		*/
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
