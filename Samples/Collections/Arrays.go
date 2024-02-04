package main

import "fmt"

/**
an array is a numbered sequence of elements of a specific length
*/
func main() {

	//By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("Emp:", a)
	/**
	set a value at an index using the array[index] = value
	syntax, and get a value with array[index]
	*/
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	//The builtin len returns the length of an array
	fmt.Println("Length:", len(a))
	//declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//Two dimensional Array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = 1 + j
		}
	}
	fmt.Println("Two Dimensional Array", twoD)
}
