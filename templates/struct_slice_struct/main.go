package main

import (
	"log"
	"os"
	"text/template"
)

type namelist struct {
	Fname string
	Lname string
}

type car struct {
	Maker string
	Model string
	Year  int16
}

type items struct {
	Names    []namelist
	Vehicles []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	name1 := namelist{
		Fname: "Clifton",
		Lname: "Avil",
	}
	name2 := namelist{
		Fname: "Aname",
		Lname: "Bname",
	}
	name3 := namelist{
		Fname: "Cname",
		Lname: "Dname",
	}
	car1 := car{
		Maker: "Mahindra",
		Model: "Thar",
		Year:  2019,
	}
	car2 := car{
		Maker: "Mahindra",
		Model: "XUV500",
		Year:  2018,
	}
	car3 := car{
		Maker: "Mahindra",
		Model: "Thar",
		Year:  2016,
	}
	car4 := car{
		"Mahindra",
		"Thar",
		2017,
	}
	AllNames := []namelist{name1, name2, name3}
	AllCars := []car{car1, car2, car3, car4}

	FinalData := items{
		Names:    AllNames,
		Vehicles: AllCars,
	}
	err := tpl.Execute(os.Stdout, FinalData)
	if err != nil {
		log.Fatalln(err)
	}
}
