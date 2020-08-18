package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Description string
	Price             float32
}

type meal struct {
	Meal  string
	Items []item
}

type menu []meal


type restaurant struct{
	Name string
	Menu menu
}
func main() {
	m := menu{
		meal{
			Meal: "Breakfast",
			Items: []item{
				item{
					Name:        "Orange Juice",
					Description: "Fresh natural orange juice.",
					Price:       50,
				},
				item{
					Name:        "Omellete",
					Description: "Egg Omelette with onions and chilly.",
					Price:       30,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Items: []item{
				item{
					Name:        "Salad",
					Description: "Freshly cut vegetables and fruits",
					Price:       50,
				},
				item{
					Name:        "Omellete",
					Description: "Egg Omelette with onions and chilly.",
					Price:       30,
				},
			},
		},
	}
	restaurants := []restaurant{
		restaurant{
			Name : "Haldiram",
			Menu : m,
		},
		restaurant{
			Name : "BTW",
			Menu : m,
		},
	}
	// log.Println(m)
	tpl := template.Must(template.ParseFiles("tmpl.html"))

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
