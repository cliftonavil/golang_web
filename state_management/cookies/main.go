package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read/", read)
	http.HandleFunc("/count/", get_count)
	http.HandleFunc("/expire/", expire)
	servererror := http.ListenAndServe(":8080", nil)
	if servererror != nil {
		log.Fatal("ListenAndServe:", servererror)
	}
}

//Set cookie
func set(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Am Writeing Cookies!!!!")
	http.SetCookie(w, &http.Cookie{
		Name:  "Clifton",
		Value: "0",
	})
	fmt.Fprintln(w, "Cookie has been written in browser!!!")
}

//Read cookie
func read(w http.ResponseWriter, r *http.Request) {
	fmt.Println("read Cookies!!!")
	c, err := r.Cookie("Clifton")
	if err != nil {
		http.Error(w, err.Error(), http.StatusSeeOther)
		return
	}

	fmt.Fprintln(w, "Cookie value is :", c)
}

//get url hitting count
func get_count(w http.ResponseWriter, r *http.Request) {
	fmt.Println("read Cookies!!!")
	cookie, err := r.Cookie("Clifton")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "Clifton",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatal(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Your total url hits : ", cookie.Value)
}

// Delete Cookies
func expire(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Cookies!!!")
	cookie, err := r.Cookie("Clifton")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	cookie.MaxAge = -1 //Force expire Cookies
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/read/", http.StatusSeeOther)
}
