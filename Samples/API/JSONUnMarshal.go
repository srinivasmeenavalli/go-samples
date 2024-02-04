package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	byt := []byte(`{"num":6.13,"strs":["pen","pencil"]}`)
	var data map[string]interface{}
	//decoding, check for associated errors
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	/*
		In order to use the values in the decoded map,
		we’ll need to convert them to their appropriate type
	*/
	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	/*
		decode JSON into custom data types.
		This has the advantages of adding
		 additional type-safety to our programs
		 and eliminating the need for
		 type assertions when accessing the decoded data
	*/
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	/*
		We can also stream JSON encodings
		 directly to os.Writers like
		 os.Stdout or even HTTP response bodies.
	*/
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
