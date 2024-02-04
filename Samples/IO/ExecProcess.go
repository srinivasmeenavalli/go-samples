package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
Sometimes we just want to completely replace
 the current Go process with another (perhaps non-Go) one.
 To do this we’ll use Go’s implementation of the classic exec function

*/
func main() {
	//Exec also needs a set of environment variables to use
	os.Setenv("PATH", `C:\Windows\System32`)
	cmd := exec.Command("systeminfo.exe")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

}
