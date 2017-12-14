package api

import (
	"github.com/aimir/goplaces/repository"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
)

func GetCountriesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repository.GetAllCountries())
}

func GetContinentCountriesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	continent, _ := strconv.Atoi(params["continent"])
	json.NewEncoder(w).Encode(repository.GetAllCountriesByContinentID(continent))
}

func GetContinentCodeCountriesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(repository.GetAllCountriesByContinentCode(params["continent"]))
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
