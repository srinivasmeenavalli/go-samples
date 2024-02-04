package main

/**
Variadic functions can be called with
any number of trailing arguments
. For example, fmt.Println is a common variadic function
*/
import "fmt"

func main() {
	/*Variadic functions can be called in the
	usual way with individual arguments
	*/
	sum2 := sum(10, 20)
	fmt.Println("Sum(10,20):", sum2)
	sum3 := sum(10, 20, 30)
	fmt.Println("Sum(10,20,30):", sum3)
	sum5 := sum(10, 20, 30, 40, 50)
	fmt.Println("Sum(10, 20, 30, 40, 50):", sum5)
	/**
	If you already have multiple args in a slice,
	apply them to a variadic function using func(slice...) like this
	*/
	intArray := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Sum 10 to 100:", sum(intArray...))
}
func sum(nums ...int) int {
	total := 0

	for num := range nums {
		total += nums[num]
	}
	return total

}
