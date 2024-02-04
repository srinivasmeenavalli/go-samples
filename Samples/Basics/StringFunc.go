package main

/*
The standard library’s strings package
provides many useful string-related functions
*/
import (
	"fmt"
	s "strings"
)

// alias fmt.Println to a shorter name
var p = fmt.Println

func main() {
	//These are functions from the strings package
	p("Contains:  ", s.Contains("test", "es"))
	p("count:  ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
