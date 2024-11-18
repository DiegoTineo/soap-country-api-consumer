package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/services"
	"github.com/gorilla/mux"
)

func GetCountryByISO(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iso := vars["iso"]
	countries, err := services.GetCountries()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	for _, country := range countries {
		if strings.EqualFold(country.IsoCode, iso) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(country)
			return
		}
	}

	http.Error(w, "Country not found", http.StatusNotFound)
}
