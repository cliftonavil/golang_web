package main

import (
	"encoding/json"
	"io/ioutil"
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

type loc struct {
	Lat float32 `json:lat`
	Log float32 `json:log`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	location := loc{}
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Something wrong ", err)
	}
	err = json.Unmarshal(jsn, &location)
	if err != nil {
		log.Fatal("Decodeing error", err)
	}

	log.Println("Recived Data", location)
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
}

func main() {

}
