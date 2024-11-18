package cli

import (
	"fmt"
	"log"

	"github.com/DiegoTineo/soap-country-api-consumer/data/gen/country_gen"
	"github.com/hooklift/gowsdl/soap"
)

func CountriesList() {

	countryServiceClient := soap.NewClient("http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso")
	countryService := country_gen.NewCountryInfoServiceSoapType(countryServiceClient)

	fullCountryInfoRequest := &country_gen.CountriesUsingCurrency{SISOCurrencyCode: "USD"}
	fullCountryInfoResponse, err := countryService.CountriesUsingCurrency(fullCountryInfoRequest)
	if err != nil {
		log.Fatalf("Error calling FullCountryInfo: %v", err)
	}

	for i := range fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName {
		fmt.Printf("Country Code: %s, Country Name: %s\n",
			fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SISOCode,
			fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SName)
	}
}