package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func (p person) SomeMethod() int {
	return 100
}
func (p person) DoubleAge() int {
	return p.Age * 2
}
func (p person) Add(addval int) int {
	return addval + 200
}
func main() {
	n := person{
		"I am Clifton",
		24,
	}
	err := tpl.Execute(os.Stdout, n)
	if err != nil {
		log.Fatalln(err)
	}
}
