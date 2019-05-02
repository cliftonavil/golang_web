package main

import "fmt"

// this function will be called
func gettotal(a, b float64) float64 {
	return a + b
}

func main() {
	//if condition
	a := 10
	b := 20

	if a > b {
		fmt.Println(a, "is greater then", b)
	} else {
		fmt.Println(b, "is greater then", a)
	}
	// passing parameters to function
	fmt.Println("function returned :", gettotal(10, 90))

}
