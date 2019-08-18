package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Fname    string
	Lname    string
	Subjects []string
}

func main() {
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		"Clifton",
		"Avil",
		[]string{"Maths", "English", "Science"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		"Clifton",
		"Avil",
		[]string{"Maths", "English", "Science"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
