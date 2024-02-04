package main

import (
	"flag"
	"fmt"
	"os"
)

/*
The flag package lets us easily define
 simple subcommands that have their own flags.
 usage:
   go run .\CmdLineSubCommands.go foo -enable -name=joe a1 a2
   go run .\CmdLineSubCommands.go bar -level a1
   go run .\CmdLineSubCommands.go bar -enable=true a1


*/
func main() {
	/*
		Declare a subcommand using the NewFlagSet function,
		and proceed to define new flags specific for this subcommand
	*/
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
