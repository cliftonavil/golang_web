package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	servererror := http.ListenAndServe(":8080", nil)
	if servererror != nil {
		log.Fatal("ListenAndServe:", servererror)
	}
	http.HandleFunc("/", foo)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	fmt.Println(r.URL)    // Get URL
	fmt.Println(r.Method) // Get Method
	io.WriteString(w, "You Search :"+v)
}
