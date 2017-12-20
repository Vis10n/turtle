package main

import (
	"log"
	"net/http"

	"turtle/crossbow"
)

func main() {
	configuration, err := ReadConfig()
	if err != nil {
		configuration.ServerPort = ":7749"
	}

	go http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))
	//router := routing.Routing()
	//
	////router.Handle("/static/", http.FileServer(http.Dir("./views/public")))
	//fmt.Println("httpdir ./views/public")
	var staticDir = "./views/public"
	http.Handle("/static/", http.FileServer(http.Dir(staticDir)))

	http.HandleFunc("/", crossbow.Weed)
	http.HandleFunc("/login", crossbow.LogIn)
	http.HandleFunc("/logout", crossbow.LogOut)

	http.HandleFunc("/weed/", crossbow.UltraWeed)
	log.Println("running server on", configuration.ServerPort)
	//log.Fatal(http.ListenAndServe(configuration.ServerPort, router))
	log.Fatal(http.ListenAndServeTLS(configuration.ServerPort, "cert.pem", "key.pem", nil))

}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target, http.StatusPermanentRedirect)

}
