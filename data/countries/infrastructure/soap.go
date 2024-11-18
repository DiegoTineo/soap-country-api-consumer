package infrastructure

import (
	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/models"
	"github.com/DiegoTineo/soap-country-api-consumer/data/gen/country_gen"
	"github.com/hooklift/gowsdl/soap"
)

const URL = "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

type countriesSoapRepo struct {
	clientSoap *soap.Client
}

func NewSoapCountriesRepo(URL string) models.CountriesRepository {
	return &countriesSoapRepo{
		clientSoap: soap.NewClient(URL),
	}
}

func (r countriesSoapRepo) GetCountries() ([] models.Country, error) {
	// service
	countryService := country_gen.NewCountryInfoServiceSoapType(r.clientSoap)

	// request
	fullCountryInfoAllCountriesRequest := &country_gen.FullCountryInfoAllCountries{}

	// response
	fullCountryInfoAllCountriesResponse, err := countryService.FullCountryInfoAllCountries(fullCountryInfoAllCountriesRequest)
	if err != nil {
		return nil, err
	}

	var countries []models.Country
	for _, country := range fullCountryInfoAllCountriesResponse.FullCountryInfoAllCountriesResult.TCountryInfo {
		countries = append(countries, models.Country{
			IsoCode: country.SISOCode,
			Name: country.SName,
			CapitalCity: country.SCapitalCity,
			CurrencyIsoCode: country.SCurrencyISOCode,
			Flag: country.SCountryFlag,
			Languages: getLanguages(country.Languages.TLanguage),
		})
	}
	return countries, nil
}

func getLanguages (TLanguage []*country_gen.TLanguage) []string{
	var languages []string = []string{}
	for _, language := range TLanguage {
		languages = append(languages, language.SName)
	}
	return languages
}

	// // service
	// countryService := country_gen.NewCountryInfoServiceSoapType(r.clientSoap)

	// // request
	// listOfCountryNamesByNameRequest := &country_gen.ListOfCountryNamesByName{}

	// // response
	// listOfCountryNamesByNameResponse, err := countryService.ListOfCountryNamesByName(listOfCountryNamesByNameRequest)
	// if err != nil {
	// 	log.Fatalf("Error calling FullCountryInfo: %v", err)
	// }

	// for i := range listOfCountryNamesByNameResponse.ListOfCountryNamesByNameResult.TCountryCodeAndName {
	// 	fmt.Printf("Country Code: %s, Country Name: %s\n",
	// 		listOfCountryNamesByNameResponse.ListOfCountryNamesByNameResult.TCountryCodeAndName[i].SISOCode,
	// 		listOfCountryNamesByNameResponse.ListOfCountryNamesByNameResult.TCountryCodeAndName[i].SName)
	// }