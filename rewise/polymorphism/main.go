package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	cankilll bool
}

func (p person) speak() {
	fmt.Println(p.fname, "Hello ***")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, "bye you might", sa.cankilll)
}

// polimorphism using interface
type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func main() {
	pp := person{
		"Clifton",
		"Avil",
	}
	pp.speak()

	sa1 := secretAgent{
		person{
			"Anil",
			"ds",
		},
		true,
	}

	sa1.speak()
	//sa1.person.speak()

	saysomething(pp) //call in interface
	saysomething(sa1)
}
