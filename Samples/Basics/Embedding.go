package main

import "fmt"

/**

Go support embedding of structs and interfaces
to express a more seamless composition of types
*/
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

/** A container embeds a base.
An embedding looks like a field without a name
*/

type container struct {
	base
	str string
}

func main() {
	//When creating structs with literals, we have to initialize the embedding explicitly
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	//full path using the embedded type name
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())
	co.base.num = 5
	//Embedding structs with methods
	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer:", d.describe())
}
