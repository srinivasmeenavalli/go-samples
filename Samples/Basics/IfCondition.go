package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}
	//an if statement without an else
	if 8%4 == 0 {
		fmt.Println("8 is devisible by Four")
	}
	/**
	A statement can precede conditionals; any variables declared
	 in this statement are available in all branches
	*/
	if num := 9; num < 0 {
		fmt.Println(" Is Negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
