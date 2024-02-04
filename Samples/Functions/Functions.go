package main

/**
Functions are central in Go, Go requires explicit returns
*/
import "fmt"

func main() {
	add := addition(43, 22)
	fmt.Println(" addition of 43,22=", add)

	sub := subtraction(43, 22)
	fmt.Println(" subtraction of 43,22=", sub)
	mul := multiplication(43, 22)
	fmt.Println(" multiplication of 43,22=", mul)

	div := division(43.0, 22.0)
	fmt.Println(" division of 43/22=", div)
}

func addition(a int, b int) int {
	return a + b
}

func subtraction(a int, b int) int {
	return a - b
}

/**
When you have multiple consecutive parameters of
the same type, you may omit the type name for the like-typed
parameters up to the final parameter that declares the type
*/
func multiplication(a, b int) int {
	return a * b
}
func division(a, b float32) float32 {
	return a / b
}
