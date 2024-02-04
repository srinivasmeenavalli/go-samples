package main

/*
Command-line arguments are a common way
to parameterize execution of programs
*/
import (
	"fmt"
	"os"
)

func main() {
	//os.Args provides access to raw command-line arguments
	argsWithProg := os.Args
	// os.Args[1:] holds the arguments to the program
	argsWithoutProg := os.Args[1:]
	fmt.Println("Arguments1 : ", argsWithProg)
	fmt.Println("Arguments2 : ", argsWithoutProg)
	//You can get individual args with normal indexing
	arg := os.Args[2]
	fmt.Println(arg)
}
