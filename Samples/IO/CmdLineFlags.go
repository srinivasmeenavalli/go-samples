package main

import (
	"flag"
	"fmt"
)

/*Go provides a flag package supporting
basic command-line flag parsing For example,
 in wc -l the -l is a command-line flag
 Usage : go build CmdLineFlags.go
         .\CmdLineFlags.exe -word=opt -numb=7 -fork -svar=flag
		 .\CmdLineFlags.exe -word=opt a1 a2 a3
*/
func main() {
	/*Declare a string flag word with a default value
	"foo" and a short description.
	This flag.String function returns a string pointer
	*/
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	/*It’s also possible to declare
	an option that uses an existing var

	declared elsewhere in the program.
	Note that we need to pass in a pointer to the flag declaration function.
	*/
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	//Once all flags are declared, call flag.Parse() to execute the command-line parsing
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)

	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}
