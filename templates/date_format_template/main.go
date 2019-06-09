package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func fDMY(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"dmy": fDMY,
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	err := tpl.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
