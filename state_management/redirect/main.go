package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/redirect1/", redirect1)
	http.HandleFunc("/redirect2/", redirect2)
	http.HandleFunc("/end/", end)
	servererror := http.ListenAndServe(":8080", nil)
	if servererror != nil {
		log.Fatal("ListenAndServe:", servererror)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page called!!!", r.Method)
	w.Header().Set("Location", "/redirect1/")
	w.WriteHeader(http.StatusSeeOther)
}

func redirect1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" redirect1 Redirected when home page called!!!", r.Method)
	w.Header().Set("Location", "/redirect2/")
	w.WriteHeader(http.StatusSeeOther)
}

func redirect2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" redirect1 Redirected when redirect1 page called!!!", r.Method)
	http.Redirect(w, r, "/end/", http.StatusSeeOther)
}
func end(w http.ResponseWriter, r *http.Request) {
	fmt.Println("final !!!", r.Method)
}
