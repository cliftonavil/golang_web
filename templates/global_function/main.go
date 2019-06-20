package main

import (
	"html/template"
	"log"
	"os"
)

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {
	Mvalue := []string{"One", "Two", "Three", "Four", "Five", "Six"}
	data := struct {
		Words []string
		Lname string
	}{
		Mvalue,
		"Clifton",
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
