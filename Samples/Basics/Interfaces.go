package main

/**
Interfaces are named collections of method signatures
To implement an interface in Go,
we just need to implement all the methods in the interface
*/

import (
	"fmt"
	"math"
)

//basic interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}
func main() {
	r := rect{width: 10, height: 6}
	c := circle{radius: 12}
	measure(r)
	measure(c)
}
