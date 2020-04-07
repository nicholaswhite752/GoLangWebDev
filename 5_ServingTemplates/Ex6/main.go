package main

import (
	"log"
	"net/http"
	"html/template"
)

var tpl *template.Template

func init(){
	//Needs template/* to do all files in templates directory
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {

	//This was my solution
	/*
	if req.Method == "GET"{
		err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
		HandleError(w, err)
	}else {
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
	}
*/

	//This solution is safer, because default is GET method
	if req.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
		return
	}

	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
