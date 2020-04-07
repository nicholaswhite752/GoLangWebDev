package main

import (
	"log"
	"os"
	"text/template"
)

type hotels struct {
	Name string
	Address   string
	City  string
	Zip string
	Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h1 := hotels{
		Name:    "Best Hotel",
		Address: "123 Main Street",
		City:    "New York City",
		Zip:     "33654",
		Region:  "Northern",
	}

	h2 := hotels{
		Name:    "Average Hotel",
		Address: "54 Johnson Drive",
		City:    "Los Angeles",
		Zip:     "15269",
		Region:  "Southern",
	}

	h3 := hotels{
		Name:    "Worst Hotel",
		Address: "8080 Bad Creek",
		City:    "Gary",
		Zip:     "74523",
		Region:  "Central",
	}

	allHotel := []hotels{h1,h2,h3}

	err := tpl.Execute(os.Stdout, allHotel)
	if err != nil {
		log.Fatalln(err)
	}
}

