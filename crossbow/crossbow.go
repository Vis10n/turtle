package crossbow

import (
	"net/http"
	"turtle/sessions"
	"turtle/views"
)

func Weed(w http.ResponseWriter, r *http.Request) {
	if sessions.IsLoggedIn(r) {
		views.ShowHome(w, sessions.GetCurrentUserName(r))
		return
	} else {
		views.ShowLogin(w)
		return
	}

}
