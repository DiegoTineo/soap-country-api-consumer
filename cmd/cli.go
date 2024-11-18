package cmd

import (
	"fmt"

	"github.com/DiegoTineo/soap-country-api-consumer/data/countries/services"
)

func RunCli() {
	countries, err := services.GetCountries()
	if err != nil {
		fmt.Println("error")
	} else {

		for _, country := range countries {
			fmt.Print("==========================================================\n")
			fmt.Printf("Country: %s\n", country.Name)
			fmt.Printf("IsoCode: %s\n", country.IsoCode)
			fmt.Printf("Capital City: %s\n", country.CapitalCity)
			fmt.Printf("Currency: %s\n", country.CurrencyIsoCode)
			fmt.Printf("Currency: %s\n", country.CurrencyIsoCode)
			fmt.Println("==========================================================")
		}

		fmt.Printf("Countries: %d\n", len(countries))
	}

}
