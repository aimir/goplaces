package api

import (
	"github.com/gorilla/mux"
	"github.com/aimir/goplaces/repository"
	"net/http"
	"strconv"
	"encoding/json"
)

func GetCountryRegionCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, _ := strconv.Atoi(params["country"])
	region, _ := strconv.Atoi(params["region"])
	json.NewEncoder(w).Encode(repository.GetAllCitiesByCountryIDByRegionID(country, region))
}

func GetCountryCodeRegionCodeCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(repository.GetAllCitiesByCountryCodeByRegionCode(params["country"], params["region"]))
}

func GetCountryCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, _ := strconv.Atoi(params["country"])
	json.NewEncoder(w).Encode(repository.GetAllCitiesByCountryID(country))
}

func GetCountryCodeCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(repository.GetAllCitiesByCountryCode(params["country"]))
}

func GetCityHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(repository.GetCity(id))
}
