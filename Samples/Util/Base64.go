package main

/*
Go provides built-in support for base64 encoding/decoding.
*/
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	//Go supports both standard and URL-compatible base64

	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	/*
		Decoding may return an error, which you can check
		if you don’t already know the input to be well-formed
	*/
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	/*
		This encodes/decodes using a URL-compatible base64 format
	*/
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
