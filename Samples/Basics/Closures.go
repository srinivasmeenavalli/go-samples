package main

/**
Go supports anonymous functions,
which can form closures. Anonymous functions
are useful when you want to define a function
inline without having to name it
*/
import "fmt"

func main() {
	seed := 10
	/**
	We call intSeq, assigning the result (a function) to nextInt.
	 This function value captures its own i value, which will be
	  updated each time we call nextInt.
	*/
	nextInt := intSeq(seed)
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println()
	newseed := 10
	/**To confirm that the state is unique to that particular function,
	create and test a new one
	*/
	nextIntSeq := intSeq(newseed)
	fmt.Println(nextIntSeq())
	fmt.Println(nextIntSeq())
	fmt.Println(nextIntSeq())

}

/**
This function intSeq returns another function,
 which we define anonymously in the body of intSeq.
 The returned function closes over the variable i to form a closure
*/

func intSeq(i int) func() int {

	return func() int {
		i += 2
		return i
	}
}
