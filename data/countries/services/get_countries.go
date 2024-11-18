package services

import (
	"github.com/DiegoTineo/soap-country-api-consumer/config"
	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/models"
)

func GetCountries() ([]models.Country, error) {
	return config.Repo.GetCountries()
}
