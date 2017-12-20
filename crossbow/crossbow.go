package crossbow

import (
	"net/http"
	"turtle/views"
	"turtle/culi"
)

func LogIn(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()

	username := r.Form["username"]
	password := r.Form["password"]

	if culi.ValidUser(username[0], password[0]) {
		views.ShowHome(w, r)
		return
	} else {
		views.ShowLoginFailed(w)
	}
}