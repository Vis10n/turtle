package crossbow

import (
	"fmt"
	"net/http"
	"strconv"
	"turtle/culi"
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

func UltraWeed(w http.ResponseWriter, r *http.Request) {
	teststr := r.URL.Path[len("/weed/"):]
	test, _ := strconv.Atoi(teststr)

	interf := culi.GetUserByID(culi.GetTeams(1)[0])
	fmt.Fprintln(w, teststr, test)
	fmt.Fprint(w, interf)

}
