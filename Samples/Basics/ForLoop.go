package main

import "fmt"

func main() {

	i := 1
	//The most basic type, with a single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("#### Out Put ####")

	//A classic initial/condition/after for loop
	for i := 4; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("#### Out Put ####")

	/**
	for without a condition will loop repeatedly
	 until you break out of the loop or return
	  from the enclosing function
	*/
	for {
		fmt.Println("Breaking For Loop")
		break
	}
	//You can also continue to the next iteration of the loop
	fmt.Println("#### Out Put ####")
	for i := 4; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
