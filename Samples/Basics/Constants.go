package main

import (
	"fmt"
	"math"
)

/*const declares a constant value,
A const statement can appear anywhere a var statement can
*/

const s string = "constant"

/**
Go supports constants of character,
string, boolean, and numeric values.
**/
func main() {

	fmt.Println(s)

	const n = 500000000
	//Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)
	/** A numeric constant has no type until
	it’s given one, such as by an explicit conversion
	*/
	fmt.Println(int64(d))
	/**
	A number can be given a type by using it in a context that
	requires one, such as a variable assignment or function
	call. For example, here math.Sin expects a float64.
	*/
	var val = math.Sin(n)
	fmt.Println(val)
}
