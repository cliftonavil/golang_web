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
}
