package main

import "fmt"

func main() {
	myarray := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(myarray); i++ {
		fmt.Println(myarray[i])
	}
}
