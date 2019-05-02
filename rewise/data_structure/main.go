package main

import "fmt"

//struct struct
type person struct {
	fname string //first small letter not availabe outside package
	lname string
	//Fname string   //first Capital letter availabe outside package
	//Lname string
}

//function
func (p person) speak() {
	fmt.Println(p.fname, "Hello ***")
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

	p1 := person{
		"Clifton",
		"Avil Dsouza",
	}
	fmt.Println(p1)
	p1.speak() // function call
}
