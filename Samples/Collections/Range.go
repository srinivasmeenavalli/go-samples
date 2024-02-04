package main

import "fmt"

/**
range iterates over elements in a variety of data structures
*/
func main() {
	sum := 0
	//range to sum the numbers in a slice
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println("Sum=", sum)
	/**
	range on arrays and slices provides both the index
	and value for each entry
	*/
	for i, num := range nums {
		if num%4 == 0 {
			fmt.Println("Divisible by 4 Index : ", i)
		}
	}
	//range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}
	//Values
	for v := range kvs {
		fmt.Println("Value:", kvs[v])
	}
	/**
	range on strings iterates over Unicode code points.
	The first value is the starting byte index of the rune and the second the rune itself
	*/
	for i, c := range "go" {
		fmt.Println("Index , Rune:", i, c)
	}

}
