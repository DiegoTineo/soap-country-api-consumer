package config

import (
	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/infrastructure"
	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/models"
)

const URL = "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

var Repo models.CountriesRepository = infrastructure.NewSoapCountriesRepo(URL)