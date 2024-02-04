package main

import "fmt"

type rect struct {
	width, height int
}

//This area method has a receiver type of *rect/pointer
func (r *rect) area() int {
	r.width = 20
	return r.width * r.height
}

//value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

/**
Go automatically handles conversion
between values and pointers for method calls.
You may want to use a pointer receiver type to avoid
 copying on method calls or to allow the method to mutate the receiving struc
*/
func main() {

	r := rect{width: 10, height: 20}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r

	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

}
