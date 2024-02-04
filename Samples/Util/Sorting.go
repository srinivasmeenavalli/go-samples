package main

import (
	"fmt"
	"sort"
)

/*
Go’s sort package implements sorting for builtins and user-defined types
*/
func main() {
	strs := []string{"d", "c", "a", "b"}
	/*Sort methods are specific to the builtin type;
	here’s an example for strings
	*/
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	floats := []float64{10, 8001.12, 8000.0000345, 398, 8001, 4, 87, 8000}
	sort.Float64s(floats)
	fmt.Println("Ints", floats)
	/*
		We can also use sort to check if a slice is already in sorted order
	*/
	s := sort.Float64sAreSorted(floats)
	fmt.Println("Sorted: ", s)
}
