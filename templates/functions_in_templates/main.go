package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type NameList struct {
	Fname string
	Lname string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	// "lc": firstThree,
}

// func firstThree(s string) string {
// 	s = strings.TrimSpace(s)
// 	s = s[:3]
// 	return s
// }
func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	name1 := NameList{
		Fname: "Clifton",
		Lname: "Avil",
	}
	name2 := NameList{
		Fname: "Aname",
		Lname: "Bname",
	}
	name3 := NameList{
		Fname: "Cname",
		Lname: "Dname",
	}
	AllNames := []NameList{name1, name2, name3}
	err := tpl.Execute(os.Stdout, AllNames)
	if err != nil {
		log.Fatalln(err)
	}
}
