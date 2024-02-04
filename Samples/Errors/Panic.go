package main

import "os"

func main() {
	/*
		use panic throughout this site to check for unexpected errors
	*/
	/*A common use of panic is to abort if a function returns an error
	 */
	_, err := os.Create("C:/GoLang/tmp/newfile/file.txt")
	if err != nil {
		panic(err)
		//panic("A Problem Occured:The system cannot find the path specified")

	}
	/*
		When first panic in main fires, the program exits without reaching the rest of the code
	*/

}
