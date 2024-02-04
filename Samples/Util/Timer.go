package main

import (
	"fmt"
	"time"
)

/*
execute Go code at some point in the future, or
repeatedly at some interval. Go’s built-in timer
and ticker features make both of these tasks easy
If you just wanted to wait, you could have used time.Sleep.
One reason a timer may be useful is that you can cancel the timer before it fires
*/
func main() {

	/*
		Timers represent a single event in the future.
		You tell the timer how long you want to wait,
		and it provides a channel that will be notified at that time.
	*/
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 Fired")

	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	/* Give the timer2 enough time to fire,
	if it ever was going to, to show it is in fact stopped
	*/
	time.Sleep(5 * time.Second)
}
