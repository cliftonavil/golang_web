package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type weatherData struct {
	LocationName string   `json:locationName`
	Weather      string   `json:weather`
	Temperature  int      `json:temperature`
	Celcius      bool     `json:celcius`
	TempForecast []int    `json:temp_forecast`
	Wind         windData `json:wind`
}

// https://mholt.github.io/json-to-go/

type windData struct {
	Direction string `json:direction`
	Speed     int    `json:speed`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	weather := weatherData{
		LocationName: "Bangalore",
		Weather:      "CLoudy",
		Temperature:  19,
		Celcius:      true,
		TempForecast: []int{31, 23, 20},
		Wind: windData{
			Direction: "South",
			Speed:     20,
		},
	}
	weatherJson, err := json.Marshal(weather)
	if err != nil {
		log.Fatal("Some Error", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(weatherJson)
}

func main() {
	http.HandleFunc("/", weatherHandler)
	http.ListenAndServe(":8000", nil)
}
