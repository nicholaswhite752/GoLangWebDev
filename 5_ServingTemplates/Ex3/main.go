package main

import (
	"io"
	"net/http"
)

//Need to do the solution code, to serve it via template
//used in next hands-on solution
func main() {
	fs := http.FileServer(http.Dir("starting-files/public"))
	http.HandleFunc("/", foo)
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

