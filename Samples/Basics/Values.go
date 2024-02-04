package main

import (
	"fmt"
	"math"
)

/*

Go has various value types including strings,
 integers, floats, booleans, etc. Here are a few basic examples.

*/
func main() {
	//Strings, which can be added together with +.cls
	fmt.Println("Go" + "Lang String")
	const str = "'Hello  Go!"
	fmt.Println("String Constant =" + str)
	const num = 24
	fmt.Println("Number =", num)
	//Integers and floats.
	fmt.Println("Number/4 =", num/4)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println("Minimum of ", math.Min(12, num))
	//Booleans, with boolean operators
	fmt.Println("true && false=", true && false)
	fmt.Println("false || true=", false || true)
	fmt.Println("!false || true=", !false || true)
	fmt.Println("false || !true=", false || !true)
	fmt.Println("!false=", !false)

}
