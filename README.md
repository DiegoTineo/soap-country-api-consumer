# soap-country-api-consumer

<img src="img.png" alt="LinkedIn"/>

## Consuming a SOAP API with Golang

Consuming a SOAP API via a WSDL file, utilizing gowsdl to generate the necessary code for service interaction.

- Xml
- WSDL file

## Generate code with gowsdl

Install gowsdl cli:
```sh
go get github.com/hooklift/gowsdl
go install github.com/hooklift/gowsdl/cmd/gowsdl@latest
```

Generate the required code for service usage:
```sh
gowsdl -p country_gen -o country_info_service_gen.go ./wsdl/CountryInfoService.xml
```