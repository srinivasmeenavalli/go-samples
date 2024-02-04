package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	filename := "config.json"
	//Some file names have extensions following a dot
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	//file’s name with the extension removed
	fmt.Println(strings.TrimSuffix(filename, ext))
	/*Rel finds a relative path between
	a base and a target. It returns
	 an error if the target cannot be made relative to base
	*/
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
