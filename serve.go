package main

import (
	"log"
	"net/http"

	"fmt"
	"turtle/scout"
)

func main() {
	configuration, err := ReadConfig()
	if err!= nil{
		configuration.ServerPort = ":7749"
	}


	router := scout.Routing()

	//router.Handle("/static/", http.FileServer(http.Dir("./views/public")))
	fmt.Println("httpdir ./views/public")
	log.Println("running server on", configuration.ServerPort)
	log.Fatal(http.ListenAndServe(configuration.ServerPort, router))
}
