package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	InternetProvider string `json:"org"`
	City             string `json:"city"`
	State            string `json:"region"`
	IP               string `json:"ip"`
}

func ExternalIP() string {
	response, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}

func main() {
	response, err := http.Get("https://ipapi.co/" + ExternalIP() + "/json/")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	var responseObject Response
	json.Unmarshal(data, &responseObject)
	fmt.Println("IP Address :", responseObject.IP)
	fmt.Println("InternetProvider :", responseObject.InternetProvider)
	fmt.Println("City :", responseObject.City)
	fmt.Println("State :", responseObject.State)

}
