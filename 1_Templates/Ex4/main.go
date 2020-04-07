package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name string
	Descrip   string
	Price  float64
}

type MealTime struct{
	WhatMeal string
	Meals []item
}

type Restaurant struct{
	MealTimes []MealTime
	Name string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	brek1 := item{
		Name:    "Eggs",
		Descrip: "Egg",
		Price:   10,
	}

	brek2 := item{
		Name:    "Pancake",
		Descrip: "Fluffy",
		Price:   15,
	}

	lunch1 := item{
		Name:    "Sandwich",
		Descrip: "Has meats",
		Price:   8,
	}

	lunch2 := item{
		Name:    "Healthy Bowl",
		Descrip: "Made with Acai",
		Price:   200,
	}

	dinner1 := item{
		Name:    "Steak",
		Descrip: "Good",
		Price:   40,
	}

	dinner2 := item{
		Name:    "Tenders",
		Descrip: "Yum",
		Price:   3,
	}

	brekMeal := MealTime{
		WhatMeal: "Breakfast",
		Meals:    []item{brek1, brek2},
	}

	lunchMeal := MealTime{
		WhatMeal: "Lunch",
		Meals:    []item{lunch1, lunch2},
	}

	dinnerMeal :=  MealTime{
		WhatMeal: "Dinner",
		Meals:    []item{dinner1, dinner2},
	}



	allFood := []MealTime{brekMeal,lunchMeal,dinnerMeal}

	r1 := Restaurant{
		MealTimes: allFood,
		Name:      "Best Food",
	}

	r2 := Restaurant{
		MealTimes: allFood,
		Name:      "Great Eats",
	}

	allRest := []Restaurant{r1,r2}

	err := tpl.Execute(os.Stdout, allRest)
	if err != nil {
		log.Fatalln(err)
	}
}


