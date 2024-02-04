package main

/**

Go supports pointers, allowing you to pass references
to values and records within your program
*/
import "fmt"

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	//The &i syntax gives the memory address of i, i.e. a pointer to i
	fmt.Println("pointer address:", &i)

}

/** zeroval will get a copy of ival
distinct from the one in the calling function
*/
func zeroval(ival int) {
	ival = 0
}

/* The *iptr code in the function body then dereferences
the pointer from its memory address to the current value at that address
*/
func zeroptr(iptr *int) {
	*iptr = 0
}
