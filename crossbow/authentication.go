package crossbow

import (
	"log"
	"net/http"
	"regexp"
	"turtle/culi"
	"turtle/sessions"
	"turtle/views"
)

var regexUsername = `^[[:alnum:]]{4,32}$`
var regexPassword = `^[[:alnum:]]{6,32}$`
var regexDept = `^(\pL|\d)((\pL)*\s*\d*)+$`
var regexName = `^((\pL)+\s*)+$`

func LogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	r.ParseForm()

	//password := r.Form["password"]
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	valid1, _ := regexp.MatchString(regexUsername, username)
	valid2, _ := regexp.MatchString(regexPassword, password)

	session, err := sessions.Store.Get(r, "session")
	if err != nil {
		log.Println("error identifying session")
		views.ShowLogin(w)
		return
	}

	if valid1 && valid2 {
		if culi.ValidUser(username, password) {
			session.Values["loggedin"] = "true"
			session.Values["username"] = username
			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusFound)
			return
		} else {
			views.ShowLoginFailed(w)
			return
		}
	} else {
		views.ShowLoginFailed(w)
		return
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.ShowSignUpForm(w)
		return
	}
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		name := r.Form.Get("name")
		dept := r.Form.Get("department")

		valid1, _ := regexp.MatchString(regexUsername, username)
		valid2, _ := regexp.MatchString(regexPassword, password)
		valid3, _ := regexp.MatchString(regexName, name)
		valid4, _ := regexp.MatchString(regexDept, dept)
		if valid1 && valid2 && valid3 && valid4 {
			err := culi.CreateUser(username, password, name, dept)
			if err != nil {
				views.ShowSignUpFailed(w)
				log.Print("Sign up username: %s failed.", username)
				return
			} else {
				views.ShowSignUpSuccess(w)
				return
			}
		} else {
			views.ShowSignUpFailed(w)
			return
		}
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "session")
	if err == nil {
		if session.Values["loggedin"] != "false" {
			session.Values["loggedin"] = "false"
			session.Save(r, w)
		}
	}
	http.Redirect(w, r, "/", 302)
}
