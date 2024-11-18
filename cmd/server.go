package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DiegoTineo/soap-country-api-consumer/presentation/api/countries/countries_routes"
	"github.com/DiegoTineo/soap-country-api-consumer/presentation/api/web/web_routes"
	"github.com/gorilla/mux"
)

func ListAndServe() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8000"
	}
	address := ":" + port

	r := mux.NewRouter()
	web_routes.RegisterHomeRoutes(r)
	countries_routes.RegisterCountriesRoutes(r)

	fmt.Println("Server running on port", port)
	http.ListenAndServe(address, r)
}
