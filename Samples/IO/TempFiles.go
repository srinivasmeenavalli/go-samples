package main

/*
Throughout program execution, we often
want to create data that isn’t needed
after the program exits. Temporary files and
directories are useful for this purpose
since they don’t pollute the file system over time
*/
import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	/*The easiest way to create a temporary
	file is by calling os.CreateTemp.
	It creates a file and opens it for reading and writing
	We provide "" as the first argument,
	so os.CreateTemp will create the file in the default location for our OS
	*/
	f, err := os.CreateTemp("", "sample")
	check(err)
	fmt.Println("Temp file name:", f.Name())
	/*Clean up the file after we’re done.
	The OS is likely to clean up temporary files
	by itself after some time, but it’s good practice to do this explicitly
	*/
	defer os.Remove(f.Name())
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	/*If we intend to write many temporary files,
	  we may prefer to create a temporary directory.
	  os.MkdirTemp’s arguments are the same as CreateTemp’s,
	  but it returns a directory name rather than an open file
	*/
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)
	defer os.RemoveAll(dname)
	/*we can synthesize temporary
	file names by prefixing them
	with our temporary directory.
	*/

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

}
