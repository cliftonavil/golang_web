package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExternalIP() string {
	response, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}

func main() {
	fmt.Println(ExternalIP())

	response, err := http.Get("https://tools.keycdn.com/geo.json?host=" + ExternalIP())
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
