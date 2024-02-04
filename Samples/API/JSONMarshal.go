package main

import (
	"encoding/json"
	"fmt"
)

/*

Go offers built-in support for JSON encoding
and decoding, including to and from built-in and custom data types
*/
type response1 struct {
	Page   int
	Fruits []string
}

/*
two structs to demonstrate encoding and decoding of custom types
Only exported fields will be encoded/decoded in JSON.
Fields must start with capital letters to be exported
*/
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	/*
		The JSON package can automatically encode your custom data types.
		It will only include exported fields in the encoded
		output and will by default use those names as the JSON keys
	*/
	res1D := &response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	res2D := &response2{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}
