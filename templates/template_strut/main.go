package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Name struct {
	Fname string
	Sname string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	myname := Name{
		Fname: "Clifton",
		Sname: "Avil",
	}

	err := tpl.Execute(os.Stdout, myname)
	if err != nil {
		log.Fatalln(err)
	}
}
