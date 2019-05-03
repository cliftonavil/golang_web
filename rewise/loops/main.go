package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	var t int = today.Day()
	switch t {
	case 0:
		fmt.Println("xxx", t)
		fallthrough
	case 1:
		fmt.Println("xxx", t)
		fallthrough
	case 3:
		fmt.Println("xxx", t)
		fallthrough
	case 4:
		fmt.Println("zzzz", t)
		fallthrough
	default:
		fmt.Println("default", t)
	}

	for k := 1; k <= 5; k++ {
		fmt.Println(k)
	}
	mymap := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range mymap {
		fmt.Println("Index :", index, " Element :", element)
	}

	for key := range mymap {
		fmt.Println(key)
	}

	for key, value := range mymap {
		fmt.Println(key, value)
	}

	mylist := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	for range mylist {
		fmt.Println(mylist)
	}
}
