package main

/*
Go has several useful functions
for working with directories in the file system
*/
import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	/*Create a new sub-directory in the
	current working directory
	*/
	err := os.Mkdir("subdir", 0755)
	check(err)
	/*When creating temporary directories,
	it’s good practice to defer their removal.
	os.RemoveAll will delete a whole directory tree (similarly to rm -rf)
	*/
	defer os.RemoveAll("subdir")
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")
	/*We can create a hierarchy of directories,
	including parents with MkdirAll.
	This is similar to the command-line mkdir -p
	*/
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")
	c, err := os.ReadDir("subdir")
	check(err)
	fmt.Println("Listing subdir")
	for _, entry := range c {
		if !entry.IsDir() {
			fmt.Println("FileName: ", entry.Name(), ": IsDirectory:", entry.IsDir())
		} else {
			fmt.Println("FolderName: ", entry.Name(), ": IsDirectory:", entry.IsDir())
		}

	}
	os.RemoveAll("subdir")
}
