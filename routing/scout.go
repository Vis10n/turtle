package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"turtle/crossbow"
)

func Routing() *mux.Router {
	var staticDir = "./views/public"
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(staticDir)))

	r.HandleFunc("/", crossbow.Weed)
	r.HandleFunc("/login", crossbow.LogIn)
	r.HandleFunc("/logout", crossbow.LogOut)

	r.HandleFunc("/weed/{teststr}", crossbow.UltraWeed)
	return r
}
