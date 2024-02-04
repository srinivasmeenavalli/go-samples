package main

/*
Writing files in Go follows similar patterns
*/

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("C:/GoLang/tmp/defer_write1.txt", d1, 0644)
	check(err)
	f, err := os.Create("C:/GoLang/tmp/defer_write2.txt")
	check(err)
	/*
		It’s idiomatic to defer a Close immediately after opening a file
	*/
	defer f.Close()
	//Write byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	//Issue a Sync to flush writes to stable storage
	f.Sync()
	//bufio provides buffered writers
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	/*
		Use Flush to ensure all buffered
		operations have been applied
		 to the underlying writer
	*/
	w.Flush()
}
