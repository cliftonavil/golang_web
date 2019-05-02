package main

import "fmt"

type animals struct {
	name   string
	number int
}

func (aa animals) call_that_animal() {
	fmt.Println("Animal name is ", aa.name, " count is", aa.number)
}

func main() {
	a := animals{
		"Cat",
		5,
	}
	a.call_that_animal()
}
