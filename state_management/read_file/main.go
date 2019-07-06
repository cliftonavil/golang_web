package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", file_upload)
	http.HandleFunc("/file/", file_operation)
	servererror := http.ListenAndServe(":8080", nil)
	if servererror != nil {
		log.Fatal("ListenAndServe:", servererror)
	}

}

func file_upload(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("index.html"))
	template.ExecuteTemplate(w, "index.html", nil)
}

func file_operation(w http.ResponseWriter, r *http.Request) {
	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		// Open FIle
		f, h, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		//Print values
		fmt.Println("\n file", f, "\n header", h, "\n error", err)

		//read file
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
		// store file in server
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
	var template = template.Must(template.ParseFiles("index.html"))
	template.ExecuteTemplate(w, "index.html", s)
}
