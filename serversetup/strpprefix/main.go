package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", swiming)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets/swiming.jpeg"))))
	servererror := http.ListenAndServe(":8080", nil)
	if servererror != nil {
		log.Fatal("ListenAndServe:", servererror)
	}
}
func swiming(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/swiming.jpeg"></img>`)
}
