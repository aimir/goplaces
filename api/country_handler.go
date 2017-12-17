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

func GetCountriesHandler(w http.ResponseWriter, r *http.Request) {
	countries := repository.GetAllCountries()
	sortCountries(countries, r)
	json.NewEncoder(w).Encode(countries)
}

func GetContinentCountriesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	continent, _ := strconv.Atoi(params["continent"])
	countries := repository.GetAllCountriesByContinentID(continent)
	sortCountries(countries, r)
	json.NewEncoder(w).Encode(countries)
}

func GetContinentCodeCountriesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	countries := repository.GetAllCountriesByContinentCode(params["continent"])
	sortCountries(countries, r)
	json.NewEncoder(w).Encode(countries)
}

func GetCountryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(repository.GetCountry(id))
}

func GetCountryByCodeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(repository.GetCountryByCode(params["code"]))
}

func sortCountries(c []*model.Country, r *http.Request) {
	queryValues := r.URL.Query()
	sortBy := queryValues.Get("sort")
	if sortBy == "" {
		sortBy = "name"
	}
	desc := strings.ToLower(queryValues.Get("direct")) == "desc"
	switch strings.ToLower(sortBy) {
	case "code":
		model.SortCountriesByCode(c, desc)
	default:
		model.SortCountriesByName(c, desc)
	}
}
