package main

import (
	"log"
	"os"
	"text/template"
)

func init() {
	tpl = template.Must(template.ParseGlob("mytemplate/*.gohtml"))
}

var tpl *template.Template

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "Mangalore")
	if err != nil {
		log.Fatalln(err)
	}
}
