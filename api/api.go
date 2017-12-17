package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

func StartAPI() error {
	r := initRouting()
	port := viper.GetInt("api_port")

	return http.ListenAndServe(fmt.Sprintf(":%d", port), middleware(r))
}

func initRouting() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	//Continents
	s.HandleFunc("/continent", GetContinentsHandler).Methods("GET")
	s.HandleFunc("/continent/{id:[0-9]+}", GetContinentHandler).Methods("GET")
	s.HandleFunc("/continent/{code:[A-Z]+}", GetContinentByCodeHandler).Methods("GET")

	//Countries
	s.HandleFunc("/country", GetCountriesHandler).Methods("GET")
	s.HandleFunc("/country/{id:[0-9]+}", GetCountryHandler).Methods("GET")
	s.HandleFunc("/country/{code:[A-Z]+}", GetCountryByCodeHandler).Methods("GET")
	s.HandleFunc("/continent/{continent:[0-9]+}/country", GetContinentCountriesHandler).Methods("GET")
	s.HandleFunc("/continent/{continent:[A-Z]+}/country", GetContinentCodeCountriesHandler).Methods("GET")

	//Regions
	s.HandleFunc("/country/{country:[0-9]+}/region", GetCountryRegionsHandler).Methods("GET")
	s.HandleFunc("/country/{country:[A-Z]+}/region", GetCountryCodeRegionsHandler).Methods("GET")
	s.HandleFunc("/region/{id:[0-9]+}", GetRegionHandler).Methods("GET")

	//Cities
	s.HandleFunc("/country/{country:[0-9]+}/city", GetCountryCitiesHandler).Methods("GET")
	s.HandleFunc("/country/{country:[A-Z]+}/city", GetCountryCodeCitiesHandler).Methods("GET")
	s.HandleFunc("/country/{country:[0-9]+}/region/{region:[0-9]+}/city", GetCountryRegionCitiesHandler).Methods("GET")
	s.HandleFunc("/country/{country:[A-Z]+}/region/{region:[A-Z0-9]+}/city", GetCountryCodeRegionCodeCitiesHandler).Methods("GET")
	s.HandleFunc("/city/{id:[0-9]+}", GetCityHandler).Methods("GET")

	return r
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", viper.GetString("api_allow_origin"))
		h.ServeHTTP(w, r)
	})
}
