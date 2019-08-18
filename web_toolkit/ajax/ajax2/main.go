package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/checkuser", checkuser)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func checkuser(w http.ResponseWriter, req *http.Request) {
	sampleUsers := map[string]bool{
		"cliftonavil@gmail.com": true,
	}

	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	fmt.Println("USERNAME: ", sbs)

	fmt.Fprint(w, sampleUsers[sbs])
}
