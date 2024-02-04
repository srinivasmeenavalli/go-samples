package main

import (
	"fmt"
	"sync"
)

/*
mutex can safely access data across multiple goroutines
to update it concurrently from multiple goroutines,  add a Mutex to synchronize access
*/
type Container struct {
	//Container holds a map of counters
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	/*
		Lock the mutex before accessing counters;
		 unlock it at the end of the function using a defer statement
	*/
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0, "c": 0},
	}
	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()

	}
	/*
		Run several goroutines concurrently;
		 note that they all access the same Container, and few of them access the same counter
	*/
	wg.Add(5)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("c", 10000)
	go doIncrement("c", 20000)
	wg.Wait()
	fmt.Println(c.counters)

}
