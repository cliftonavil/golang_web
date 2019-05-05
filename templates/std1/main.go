package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, error := template.ParseFiles("tpl.gohtml")
	if error != nil {
		log.Fatalln(error)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()
	error = tpl.Execute(os.Stdout, nil)
	if error != nil {
		log.Fatalln(error)
	}
}
