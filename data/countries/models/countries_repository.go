package models

type CountriesRepository interface {
	GetCountries() ([]Country, error)
}