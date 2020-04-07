package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	//in index.html
	//<img src="/resources/pics/dog.jpeg">
	//so it strips the /resources
	//Then looks for /pics/dog.jpeg starting in the public directory
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
