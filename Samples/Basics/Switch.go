package main

import (
	"fmt"
	"time"
)

/**
Switch statements express conditionals across many branches.
*/
func main() {
	//Here’s a basic switch.
	i := 2
	fmt.Println("Write", i, "as ")
	switch i {
	case 1:
		fmt.Println("Case One")
	case 2:
		fmt.Println("Case Two")
	case 3:
		fmt.Println("Case Three")

	}
	/**
	You can use commas to separate multiple expressions
	 in the same case statement. We use the optional default
	 case in this example as well
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Case: It's a Weekend")
	default:
		fmt.Println("Case: It's a Week day")
	}
	/** switch without an expression is an
	alternate way to express if/else logic
	*/
	t := time.Now().Hour()
	fmt.Println("Current Time= ", t)
	switch {
	case t < 12:
		fmt.Println("Case: Good Morning ")
	case t >= 12 && t < 18:
		fmt.Println("Case: Good Afternoon ")
	case t >= 18 && t <= 21:
		fmt.Println("Case: Good Evening ")
	default:
		fmt.Println("Case: Good Night")

	}
	/**
	A type switch compares types instead of values
	. You can use this to discover the type of an interface value
	*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
