package main

import (
	"html/template"
	"log"
	"os"
)

type Employee struct {
	Fname, Lname, City string
	Admin              bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {
	e1 := Employee{
		Fname: "Clifton",
		Lname: "Dsouza",
		City:  "Bangalore",
		Admin: true,
	}
	e2 := Employee{
		Fname: "Austin",
		Lname: "Dsouza",
		City:  "Mangalore",
		Admin: false,
	}
	e3 := Employee{
		Fname: "Sunil",
		Lname: "Dsouza",
		City:  "Pune",
		Admin: true,
	}
	e4 := Employee{
		Fname: "Anil",
		Lname: "Kumar",
		City:  "Bangalore",
		Admin: true,
	}
	data := []Employee{e1, e2, e3, e4}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
