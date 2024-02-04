package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
Go programs can intelligently handle Unix signals.For example
server to gracefully shutdown when it receives a SIGTERM,
or a command-line tool to stop processing input if it receives a SIGINT

*/
func main() {
	/*Go signal notification works
	by sending os.Signal values on a channel*/
	sigs := make(chan os.Signal, 1)
	/*signal.Notify registers the
	given channel to receive notifications of the specified signals
	*/
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	/*This goroutine executes a blocking receive
	for signals. When it gets one it’ll print it out
	and then notify the program that it can finish
	*/
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true

	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}

/*Note:  When we run this program it will block waiting for a signal.
  By typing ctrl-C (which the terminal shows as ^C)
  we can send a SIGINT signal,
  causing the program to print interrupt and then exit.
*/
