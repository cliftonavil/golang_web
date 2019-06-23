package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", img1)
	http.HandleFunc("/swiming.jpeg", img2)
	servererror := http.ListenAndServe(":8080", nil)
	if servererror != nil {
		log.Fatal("ListenAndServe:", servererror)
	}
}
func img1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="swiming.jpeg"></img>`)
}
func img2(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("swiming.jpeg")
	if err != nil {
		http.Error(w, "Image Not Found", 404)
		return
	}
	defer f.Close()
	// io.Copy(w, f)

	fi, ere := f.Stat()
	if ere != nil {
		http.Error(w, "Content Not Found", 404)
		return
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
