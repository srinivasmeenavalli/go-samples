package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	os.Setenv("PATH", `C:\Windows\System32`)
	/*The exec.Command helper creates
	an object to represent this external process
	*/
	pingCmd := exec.Command("ping.exe", "Google.com")
	/*
		.Output is another helper that handles
		the common case of running a command,
		waiting for it to finish, and collecting its output
	*/
	pingOut, err := pingCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> Ping DNS")
	fmt.Println(string(pingOut))
	/*
		Here we explicitly grab input/output pipes,
		start the process, write some input to it,
		read the resulting output, and finally wait for the process to exit.
	*/
	fmt.Println("> System Configurations")
	ipConfigCmd := exec.Command("ipconfig.exe")
	ipConfigIn, _ := ipConfigCmd.StdinPipe()
	ipConfigOut, _ := ipConfigCmd.StdoutPipe()
	ipConfigCmd.Start()
	ipConfigIn.Write([]byte("hello grep\ngoodbye grep"))
	ipConfigIn.Close()
	ipConfigBytes, _ := io.ReadAll(ipConfigOut)
	ipConfigCmd.Wait()

	fmt.Println("> Ipconfig Command hello")
	fmt.Println(string(ipConfigBytes))

	sysInfoCmd := exec.Command("systeminfo.exe")
	sysInfoOut, err := sysInfoCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> systeminfo.exe")
	fmt.Println(string(sysInfoOut))

}
