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
	err := tpl.Execute(os.Stdout, Mvalue)
	if err != nil {
		log.Fatalln(err)
	}
}
