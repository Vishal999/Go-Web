package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Name   string
	Number string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcademicYear string
	Fall         semester
	Spring       semester
	Summer       semester
}

func main() {

	tpl := template.Must(template.ParseFiles("tmpl.gohtml"))
	years := []year{
		year{
			AcademicYear: "2020-2021",
			Fall: semester{
				Term: "FALL",
				Courses: []course{
					course{
						Name:   "Web Programming",
						Number: "WP101",
						Units:  "10",
					},
					course{
						Name:   "Intro To GO",
						Number: "GO101",
						Units:  "8",
					},
				},
			},
			Spring: semester{
				Term: "SPRING",
				Courses: []course{
					course{
						Name:   "Web-Design",
						Number: "WD101",
						Units:  "10",
					},
					course{
						Name:   "Advance GO",
						Number: "GO102",
						Units:  "8",
					},
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}
}
