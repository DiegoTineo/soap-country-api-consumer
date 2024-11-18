package models

type Country struct {
	IsoCode         string
	Name            string
	CapitalCity     string
	Languages       []string
	CurrencyIsoCode string
	Flag            string
}