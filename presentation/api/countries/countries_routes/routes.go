package countries_routes

import (
	"github.com/DiegoTineo/soap-country-api-consumer/presentation/api/countries/handlers"
	"github.com/gorilla/mux"
)

func RegisterCountriesRoutes(r *mux.Router) {
	s := r.PathPrefix("/api/countries").Subrouter()
	s.HandleFunc("", handlers.GetCountries).Methods("GET")
	s.HandleFunc("/{iso}", handlers.GetCountryByISO).Methods("GET")
}
