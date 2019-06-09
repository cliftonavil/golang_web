package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func Fdouble(f int) int {
	return f + f
}

func Fsquare(f int) float64 {
	return math.Pow(float64(f), 2)
}

func FSquareRoot(f int) float64 {
	return math.Sqrt(float64(f))
}

var fm = template.FuncMap{
	"Fdouble":     Fdouble,
	"Fsquare":     Fsquare,
	"FSquareRoot": FSquareRoot,
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	Mvalue := 10
	err := tpl.Execute(os.Stdout, Mvalue)
	if err != nil {
		log.Fatalln(err)
	}
}
