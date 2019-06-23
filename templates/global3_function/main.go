package main

import (
	"log"
	"os"
	"text/template"
)

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {
	// Compare values in template
	g1 := struct {
		Score1, Score2 int
	}{
		2,
		6,
	}
	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
