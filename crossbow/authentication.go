package crossbow

import (
	"log"
	"net/http"
	"turtle/culi"
	"turtle/sessions"
	"turtle/views"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "session")
	if err != nil {
		log.Println("error identifying session")
		views.ShowLogin(w)
		return
	}

	r.ParseForm()
	//password := r.Form["password"]
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	//TODO: Validate form

	if culi.ValidUser(username, password) {
		session.Values["loggedin"] = "true"
		session.Values["username"] = username
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
		return
	} else {
		views.ShowLoginFailed(w)
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

		//TODO: Validate form

		err := culi.CreateUser(username, password, name, dept)
		if err != nil {
			views.ShowSignUpFailed(w)
			log.Print("Sign up username: %s failed.", username)
			return
		} else {
			views.ShowSignUpSuccess(w)
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
