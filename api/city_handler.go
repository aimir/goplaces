package api

import (
	"encoding/json"
	"github.com/aimir/goplaces/model"
	"github.com/aimir/goplaces/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

//Handle request to get cities by country ID and region ID
func GetCountryRegionCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, _ := strconv.Atoi(params["country"])
	region, _ := strconv.Atoi(params["region"])
	cities := repository.GetAllCitiesByCountryIDByRegionID(country, region)
	sortCities(cities, r)
	json.NewEncoder(w).Encode(cities)
}

//Handle request to get cities by country code and region code
func GetCountryCodeRegionCodeCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cities := repository.GetAllCitiesByCountryCodeByRegionCode(params["country"], params["region"])
	sortCities(cities, r)
	json.NewEncoder(w).Encode(cities)
}

//Handle request to get cities by country ID
func GetCountryCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, _ := strconv.Atoi(params["country"])
	cities := repository.GetAllCitiesByCountryID(country)
	sortCities(cities, r)
	json.NewEncoder(w).Encode(cities)
}

//Handle request to get cities by country code
func GetCountryCodeCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cities := repository.GetAllCitiesByCountryCode(params["country"])
	sortCities(cities, r)
	json.NewEncoder(w).Encode(cities)
}

//Get city full data
func GetCityHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(repository.GetCityView(id))
}

func sortCities(c []*model.City, r *http.Request) {
	queryValues := r.URL.Query()
	desc := strings.ToLower(queryValues.Get("direct")) == "desc"
	model.SortCitiesByName(c, desc)
}
