package main

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

	err := os.RemoveAll("./subdir")
	check(err)
	err = os.Mkdir("subdir", 0755)
	check(err)

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file2.1")
	createEmptyFile("subdir/parent/file2.3")
	check(err)
	//err = os.Chdir("subdir/parent")
	//check(err)
	/*ReadDir lists directory contents,
	returning a slice of os.DirEntry objects
	*/
	c, err := os.ReadDir("subdir/parent")
	check(err)
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)

	os.RemoveAll("subdir")

}
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
