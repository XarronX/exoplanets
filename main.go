package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/million_dollar_space_programme/exoplanets/configs"
	"github.com/million_dollar_space_programme/exoplanets/routes"
)

func init() {
	flag.StringVar(&configs.DATABASE_URL, "dbconn", "", "database url")
}

func main() {
	flag.Parse()
	if configs.DATABASE_URL == "" {
		log.Fatal("no database url provided")
	}

	r := mux.NewRouter()
	routes.SetRoutes(r)

	log.Println("starting server at 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
