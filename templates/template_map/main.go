package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	marklist := map[string]int{"English": 67, "Maths": 55, "Physis": 60, "Chemistry": 65, "Hindi": 44}
	err := tpl.Execute(os.Stdout, marklist)
	if err != nil {
		log.Fatalln(err)
	}
}
