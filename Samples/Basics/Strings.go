package main

import (
	"fmt"
	"unicode/utf8"
)

/**
A Go string is a read-only slice of bytes.
The language and the standard library treat strings
 - as containers of text encoded in UTF-8
 In Go, the concept of a character is called a rune
*/
func main() {
	/*“hello” in the Thai language
	Go string literals are UTF-8 encoded text*/

	const str = "สวัสดี"
	strlen := len(str)
	/*Strings are equivalent to []byte,
	this will produce the length of the raw bytes stored within */
	fmt.Println("Length of String:", strlen)

	//Indexing into a string produces the raw byte values at each index
	for i := 0; i < strlen; i++ {
		fmt.Printf("%x ", str[i])

	}
	fmt.Println()
	/**run-time of RuneCountInString dependes on the size of the string
	, because it has to decode each UTF-8 rune sequentially
	*/
	fmt.Println("Rune Count:", utf8.RuneCountInString(str))

	/**
	A range loop handles strings specially and decodes
	each rune along with its offset in the string
	*/
	for idx, runeValue := range str {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println()
	for i, w := 0, 0; i < strlen; i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

	}
}
