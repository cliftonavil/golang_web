package main

import "fmt"

var my_age int //By default will be 0

func main() {
	x := "CLifton"
	age := 10
	fmt.Println("Age from Package level variabke :", my_age)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", age)
}
