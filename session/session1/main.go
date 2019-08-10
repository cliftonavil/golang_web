package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	Username string
	Password []byte
	Fname    string
	Lname    string
	Role     string
}

var tpl *template.Template
var dbUsers = map[string]user{}     //user ID, user
var dbSession = map[string]string{} // Session ID user ID

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["cliftonavil@gmail.com"] = user{"cliftonavil@gmail.com", bs, "Clifton", "Avil", "user"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/display", display)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	servererror := http.ListenAndServe(":8080", nil)
	if servererror != nil {
		log.Fatal("ListenAndServe:", servererror)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.html", u)
}

func display(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin" {
		http.Error(w, "Only Admin has Rights!!!", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "display.html", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		role := req.FormValue("role")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSession[c.Value] = un

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error!!", http.StatusInternalServerError)
		}

		// store user in dbUsers
		u := user{un, bs, f, l, role}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Inside")
	// if !alreadyLoggedIn(r) {
	// 	fmt.Println("error")
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// }

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		// check user registered!!!
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username does not exisits!!", http.StatusForbidden)
			return
		}
		// Check password matches
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSession[c.Value] = un
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	// delete the session
	delete(dbSession, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
