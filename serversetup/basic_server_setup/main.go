package main

import (
	"log"
	"net/http"
)

func main() {
	servererror := http.ListenAndServe(":8080", nil)

	if servererror != nil {
		log.Fatal("ListenAndServe:", servererror)
	}
}
