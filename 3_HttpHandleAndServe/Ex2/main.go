package main

import (
	"log"
	"net/http"
	"text/template"
)


func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar/", bar)
	http.HandleFunc("/user/", myName)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("error parsing template in foo", err)
	}

	err = tpl.ExecuteTemplate(w, "tpl.gohtml", "foo ran")
	if err != nil {
		log.Fatalln("error executing template in foo", err)
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("error parsing template in bar", err)
	}

	err = tpl.ExecuteTemplate(w, "tpl.gohtml", "bar ran")
	if err != nil {
		log.Fatalln("error executing template in bar", err)
	}
}

func myName(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("error parsing template in myName", err)
	}

	err = tpl.ExecuteTemplate(w, "tpl.gohtml", "User")
	if err != nil {
		log.Fatalln("error executing template in myName", err)
	}
}
