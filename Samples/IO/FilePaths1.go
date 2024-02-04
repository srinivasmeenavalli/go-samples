package main

import (
	"fmt"
	"path/filepath"
)

/*The filepath package provides functions
to parse and construct file paths in a way
that is portable between operating systems;
dir/file on Linux vs. dir\file on Windows
*/
func main() {
	/*Join should be used to construct paths
	in a portable way. It takes any number of arguments
	 and constructs a hierarchical path from them
	*/
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	/* use Join instead of concatenating
	/s or \s manually. In addition
	to providing portability,
	Join will also normalize paths by
	removing superfluous separators and directory changes
	*/
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir2", "filename"))
	/*Dir and Base can be used to split
	a path to the directory and the file.
	Alternatively, Split will return both in the same call
	*/
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	dir, file := filepath.Split(p)
	fmt.Println("Split(p)-Dir:", dir, ":File Name:", file)
	//check whether a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

}
