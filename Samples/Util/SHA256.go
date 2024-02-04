package main

/*

SHA256 hashes are frequently used to compute
short identities for binary or text blobs.
For example, TLS/SSL certificates
use SHA256 to compute a certificate’s signature
*/
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	//start with a new hash
	h := sha256.New()
	h.Write([]byte(s))
	/*
		This gets the finalized hash result
		as a byte slice. The argument to Sum
		can be used to append to an existing byte slice:
		it usually isn’t needed
	*/
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("sha256 : %x\n", bs)
}
