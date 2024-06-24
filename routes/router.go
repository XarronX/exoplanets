package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/million_dollar_space_programme/exoplanets/handlers"
)

func SetRoutes(r *mux.Router) {
	r.Handle("/exoplanet", http.HandlerFunc(handlers.CreateExoPlanet)).Methods("POST")                     // Add a new exoplanet
	r.Handle("/exoplanets", http.HandlerFunc(handlers.GetExoPlanets)).Methods("GET")                       // List all exoplanets
	r.Handle("/exoplanet/{id}", http.HandlerFunc(handlers.GetExoPlanet)).Methods("GET")                    // Get exoplanet by id
	r.Handle("/exoplanet/{id}", http.HandlerFunc(handlers.UpdateExoPlanet)).Methods("PUT")                 // Update exoplanet using id
	r.Handle("/exoplanet/{id}", http.HandlerFunc(handlers.DeleteExoPlanet)).Methods("DELETE")              // Delete exoplanet using id
	r.Handle("/exoplanet/{id}/fuel-estimation", http.HandlerFunc(handlers.FuelEstimation)).Methods("POST") // fuel-estimation for exoplanet
}
