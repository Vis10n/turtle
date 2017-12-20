package main

import (
	"log"
	"net/http"

	"fmt"
	"turtle/routing"
)

func main() {
	configuration, err := ReadConfig()
	if err != nil {
		configuration.ServerPort = ":7749"
	}

	router := routing.Routing()

	//router.Handle("/static/", http.FileServer(http.Dir("./views/public")))
	fmt.Println("httpdir ./views/public")
	log.Println("running server on", configuration.ServerPort)
	log.Fatal(http.ListenAndServe(configuration.ServerPort, router))
}
