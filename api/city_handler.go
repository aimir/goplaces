package api

import (
	"github.com/gorilla/mux"
	"github.com/aimir/goplaces/repository"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/aimir/goplaces/model"
	"strings"
)

func GetCountryRegionCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, _ := strconv.Atoi(params["country"])
	region, _ := strconv.Atoi(params["region"])
	cities := repository.GetAllCitiesByCountryIDByRegionID(country, region)
	sortCities(cities, r)
	json.NewEncoder(w).Encode(cities)
}

func GetCountryCodeRegionCodeCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cities := repository.GetAllCitiesByCountryCodeByRegionCode(params["country"], params["region"])
	sortCities(cities, r)
	json.NewEncoder(w).Encode(cities)
}

func GetCountryCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, _ := strconv.Atoi(params["country"])
	cities := repository.GetAllCitiesByCountryID(country)
	sortCities(cities, r)
	json.NewEncoder(w).Encode(cities)
}

func GetCountryCodeCitiesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cities := repository.GetAllCitiesByCountryCode(params["country"])
	sortCities(cities, r)
	json.NewEncoder(w).Encode(cities)
}

func GetCityHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(repository.GetCity(id))
}

func sortCities(c []*model.City, r *http.Request) {
	queryValues := r.URL.Query()
	desc := strings.ToLower(queryValues.Get("direct")) == "desc"
	model.SortCitiesByName(c, desc)
}
