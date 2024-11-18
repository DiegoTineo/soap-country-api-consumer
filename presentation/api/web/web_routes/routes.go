package web_routes

import (
	"net/http"

	"github.com/DiegoTineo/soap-country-api-consumer/presentation/api/web/handlers"
	"github.com/gorilla/mux"
)

func FileServer(r *mux.Router) {
	// Servir archivos estáticos desde el directorio correcto
	fs := http.FileServer(http.Dir("./presentation/api/web"))
	r.PathPrefix("").Handler(http.StripPrefix("", fs))
}

func RegisterHomeRoutes(r *mux.Router) {
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("presentation/api/web/public"))))
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
}
