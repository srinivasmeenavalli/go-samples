package main

import "fmt"

type customError struct {
	arg  int
	prob string
}

func (e *customError) Error() string {
	return fmt.Sprintf("Custom Error: %d - %s", e.arg, e.prob)
}

/**
use &argError syntax to build a new struct,
 supplying values for the two fields arg and prob
*/
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &customError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
func main() {
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("Range f2 failed:", e)
		} else {
			fmt.Println("Range f2 worked:", r)
		}

		_, error := f2(42)
		/**
		If you want to programmatically use
		the data in a custom error, you’ll
		need to get the error as an instance of
		 the custom error type via type assertion
		*/
		if ae, ok := error.(*customError); ok {
			fmt.Println("Argument:", ae.arg)
			fmt.Println("Problem:", ae.prob)
		}

	}

}
