package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/models"
	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/services"
)

type Filter struct {
	Page           int `json:"page"`
	RecordsPerPage int `json:"recordsPerPage"`
}

type Filtered struct {
	Countries []models.Country
	Page int 
  TotalPages int
}

func GetCountriesPerPages(filter Filter, countries []models.Country) []models.Country {
	start := filter.Page * filter.RecordsPerPage
	end := start + filter.RecordsPerPage

	if filter.RecordsPerPage == 0{
		end = len(countries) - 1
	}
	if start > len(countries) {
		return []models.Country{}
	}
	if end > len(countries) {
		end = len(countries)
	}
	
	return countries[start:end]
}

func GetCountries(w http.ResponseWriter, r *http.Request) {
	var filter Filter

	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		filter = Filter{
			Page:           0,
			RecordsPerPage: 0,
		}
	}

	countries, err := services.GetCountries()
	if err != nil {
		http.Error(w, "Unable to get countries", http.StatusInternalServerError)
		return
	}

	filteredCountries := GetCountriesPerPages(filter, countries)

	// Establecer el tipo de contenido a JSON
	w.Header().Set("Content-Type", "application/json")

	// Codificar los datos en formato JSON y escribirlos en la respuesta
	if err := json.NewEncoder(w).Encode(filteredCountries); err != nil {
		http.Error(w, "Unable to encode countries", http.StatusInternalServerError)
	}
}
