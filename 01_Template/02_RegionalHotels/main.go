package main

import (
	"log"
	"os"
	"text/template"
)

type Hotel struct {
	Name    string
	Address address
}

type address struct {
	City   string
	ZIP    string
	Region string
}

const (
	central string = "Central"
	south   string = "South"
	north   string = "North"
	west    string = "West"
	east    string = "East"
)

func main() {
	hotels := []Hotel{
		Hotel{
			Name: "ITC GrandChola",
			Address: address{
				City:   "New Delhi",
				ZIP:    "111001",
				Region: central,
			},
		},
		Hotel{
			Name: "Taj",
			Address: address{
				City:   "Mumbai",
				ZIP:    "990090",
				Region: south,
			},
		},
	}

	tpl := template.Must(template.ParseFiles("tmpl.html"))
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
