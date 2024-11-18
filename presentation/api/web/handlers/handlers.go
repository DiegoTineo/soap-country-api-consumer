package handlers

import (
	"html/template"
	"net/http"

	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/models"
	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("presentation/api/web/templates/index.html")
	if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
	}

	countries, err := services.GetCountries()
	if err != nil {
			countries = []models.Country{}
	}

	data := struct {
			Title     string
			Header    string
			Countries []models.Country
			Footer    string
	}{
			Title:     "Countries List",
			Header:    "Countries",
			Countries: countries,
			Footer:    "Golang is cool",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
	}
}
