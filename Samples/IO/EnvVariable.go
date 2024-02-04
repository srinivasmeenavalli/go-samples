package main

/*

Environment variables are a universal mechanism
for conveying configuration information to programs
*/
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		To set a key/value pair, use os.Setenv.
		To get a value for a key, use os.Getenv.
		 This will return an empty string
		 if the key isn’t present in the environment
	*/
	os.Setenv("FOO", "Environment Varaible Value: 1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println()
	/*Use os.Environ to list all key/value pairs
	in the environment. This returns a slice of
	 strings in the form KEY=value.
	 You can strings.SplitN them to
	 get the key and value. Here we print all the keys
	*/
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		fmt.Println(pair[0], ":", pair[1])
	}

}
