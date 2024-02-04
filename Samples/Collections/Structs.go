package main

import "fmt"

/**

Go’s structs are typed collections of fields.
They’re useful for grouping data together
to form records. Structs are mutable
*/
type person struct {
	// person struct type has name and age fields
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	/**You can safely return a pointer to local variable
	as a local variable will survive the scope of the function */
	return &p
}

func main() {
	//name the fields when initializing a struct
	person1 := person{"Bob", 20}
	fmt.Println(person1.name, person1.age)
	fmt.Println(person{"Srinivas", 43})
	fmt.Println(person{name: "RGV", age: 32})
	//Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})
	//An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))
	s := person{name: "Sean", age: 50}
	//Access struct fields with a dot
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
