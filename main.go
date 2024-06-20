package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/million_dollar_space_programme/exoplanets/routes"
)

func main() {
	r := mux.NewRouter()
	routes.SetRoutes(r)

	log.Println("starting server at 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
