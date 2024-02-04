package main

/**

Go supports recursive functions.
*/

import "fmt"

/**
This fact function calls itself until it
reaches the base case of fact(0)
*/
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fcatVal := fact(5)
	fmt.Println("Factorial of 5:", fcatVal)
	fmt.Println("Factorial of 6:", fact(6))
	/**
	Closures can also be recursive,
	but this requires the closure to be declared with
	a typed var explicitly before it’s defined.
	*/
	var fib func(c int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)

	}

	fmt.Println(fib(40))
}
