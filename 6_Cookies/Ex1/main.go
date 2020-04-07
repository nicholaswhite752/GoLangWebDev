package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie{
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path: "/",
		}
	}else{
		count, err := strconv.Atoi(c.Value)
		if err != nil {
			log.Fatalln(err)
		}
		count++
		c.Value = strconv.Itoa(count)
	}

	http.SetCookie(w, c)

}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
