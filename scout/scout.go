package scout

import (
	"turtle/views"
	"turtle/crossbow"
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router{
	var staticDir = "./views/public"
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(staticDir)))

	r.HandleFunc("/", views.ShowLogin)

	r.HandleFunc("/login", crossbow.LogIn)


	return r
}