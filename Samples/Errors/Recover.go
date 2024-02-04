package main

import "fmt"

/*
Go makes it possible to recover from a panic,
by using the recover built-in function.
 A recover can stop a panic from aborting the program
 and let it continue with execution instead
*/

func mayPanic() {
	//function panics
	panic("Problem Occured")
}
func main() {
	defer func() {
		r := recover()
		if r != nil {
			/*
				The return value of recover is the error
				raised in the call to panic
			*/
			fmt.Println("Recovered. Error:\n", r)
		}

	}()
	/*
		When the enclosing function panics,
		the defer will activate and a recover
		 call within it will catch the panic.
	*/
	mayPanic()
	fmt.Println("After mayPanic()")
}
