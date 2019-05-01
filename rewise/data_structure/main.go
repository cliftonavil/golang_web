package main

import "fmt"

//struct struct
type Hello struct {
}

func main() {
	//slice
	a := []int{1, 2, 3, 4, 5, 5, 6, 6}
	b := []string{"Cat", "Dog", "Cow"}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a) //print type
	fmt.Println(b[1])     //using index position
	//map
	age := map[string]int{
		"Anil":  20,
		"Sunil": 22,
		"Avil":  23,
		"Avin":  30,
	}
	fmt.Println(age)
	fmt.Println(age["Sunil"]) // using key get value
}
