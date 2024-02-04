package main

/**
In Go it’s idiomatic to communicate errors via an explicit,
separate return value
Go’s approach makes it easy to see which functions return errors
and to handle them
*/
import (
	"errors"
	"fmt"
)

/**By convention, errors are the last
return value and have type error, a built-in interface
*/

func f1(arg int) (int, error) {
	if arg == 42 {
		//errors.New constructs a basic error value with the given error message
		return -1, errors.New("can't work with 42")
	}
	//A nil value in the error position indicates that there was no error
	return arg + 3, nil
}
func main() {

	for _, i := range []int{7, 42, 43, 7665, 546446} {
		var r, e = f1(i)
		if e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

}
