package main

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

}
