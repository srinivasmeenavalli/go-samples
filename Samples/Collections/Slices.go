package main

import "fmt"

/**
Slices are a key data type in Go, giving a
more powerful interface to sequences than arrays.
slices support several more that make them richer than arrays
*/
func main() {

	s := make([]string, 3)
	fmt.Println("Emp:", s)
	//set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[0])
	fmt.Println("Length:", len(s))
	//append returns a slice containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)
	//Slices Slices can also be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied:", s)
	//slice of the elements from,to
	l := c[2:5]
	fmt.Println("Slice from 2:5", l)

	//slices up to but excluding
	//l = c[:15] panic: runtime error: slice bounds out of range
	//[:15] with capacity 6

	l = c[:4]
	fmt.Println("Slice upto :4 ", l)
	//slices up from (and including)
	l = c[2:]
	fmt.Println("Slice from :2 Including ", l)
	// declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	/**
	Slices can be composed into multi-dimensional data structures.
	The length of the inner slices can vary,
	unlike with multi-dimensional arrays.
	*/
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
