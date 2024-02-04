package main

/**
Go has built-in support for multiple return values.
 This feature is used often in idiomatic Go,
 for example to return both result and error values from a function
*/

import "fmt"

func main() {
	//2 different return values from the call with multiple assignment
	r1, r2 := vals()
	fmt.Println("Returned Value1,Value 2:", r1, r2)
	/**If you only want a subset of the returned values,
	use the blank identifier "_"
	*/
	r3, _ := vals()
	fmt.Println("Returned Value 1:", r3)

	_, r4 := vals()
	fmt.Println("Returned Value 2:", r4)

}

/**
The (int, int) in this function signature shows that the function returns 2 ints
*/
func vals() (int, int) {
	return 3, 34
}
