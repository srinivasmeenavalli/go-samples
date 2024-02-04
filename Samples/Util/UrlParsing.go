package main

/*
URLs provide a uniform way to locate resources.
*/

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	//Parse the URL
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	//		Accessing the scheme is straightforward.

	fmt.Println("Scheme: ", u.Scheme)
	/*
		User contains all authentication info;
		call Username and Password on this for individual values
	*/
	fmt.Println("User: ", u.User)
	fmt.Println("USer Name: ", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("Password: ", p)
	/*
		The Host contains both the hostname and the port,
		 if present. Use SplitHostPort to extract them
	*/
	fmt.Println("Host: ", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Host: ", host)
	fmt.Println("Port: ", port)
	//extract the path and the fragment after the #.
	fmt.Println("URL Path: ", u.Path)
	fmt.Println("URL Fragment: ", u.Fragment)
	//To get query params in a string of k=v format, use RawQuery
	fmt.Println("Raw Query: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("Query Map: ", m)
	fmt.Println("Query Param Value: ", m["k"][0])
}
