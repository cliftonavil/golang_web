package main

import (
	"fmt"
)

// this function will be called
func gettotal(a, b float64) float64 {
	return a + b
}

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}

func main() {
	// //if condition
	// a := 10
	// b := 20

	// if a > b {
	// 	fmt.Println(a, "is greater then", b)
	// } else {
	// 	fmt.Println(b, "is greater then", a)
	// }
	// // passing parameters to function
	// fmt.Println("function returned :", gettotal(10, 90))
	// //switch case
	// today := time.Now()
	// switch today.Day() {
	// case 5:
	// 	fmt.Println("Today is 5th. Clean your house.")
	// case 10:
	// 	fmt.Println("Today is 10th. Buy some wine.")
	// case 15:
	// 	fmt.Println("Today is 15th. Visit a doctor.")
	// case 25:
	// 	fmt.Println("Today is 25th. Buy some food.")
	// case 31:
	// 	fmt.Println("Party tonight.")
	// default:
	// 	fmt.Println("No information available for that day.")
	// }

	defer second()
	first()

}
