package main

import (
	"fmt"
	"net/http"
)

/*A fundamental concept in net/http servers is handlers.
A handler is an object implementing the http.Handler interface
*/
func hello(w http.ResponseWriter, req *http.Request) {
	/*Functions serving as handlers take
	a http.ResponseWriter and a http.Request as arguments.
	 The response writer is used to fill in the HTTP response
	*/
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	/*reading all the HTTP request headers
	and echoing them into the response body*/
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

/*HTTP basic Server */
func main() {
	/* register our handlers on server routes
	using the http.HandleFunc convenience function.
	It sets up the default router in the net/http package
	and takes a function as an argument.
	*/
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/", hello)
	/*Finally, we call the ListenAndServe
	with the port and a handler.
	nil tells it to use the default router we’ve just set up
	*/
	http.ListenAndServe(":8090", nil)
}
