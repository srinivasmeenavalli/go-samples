package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
sync/atomic package for atomic counters accessed by multiple goroutines
*/

func main() {
	/*unsigned integer to represent
	our (always-positive) counter
	*/
	var ops uint64
	/*
		WaitGroup will help us wait
		for all goroutines to finish their work
	*/
	var wg sync.WaitGroup

	/*
		50 goroutines that each increment the counter exactly 2000 times
	*/

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 2)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}
