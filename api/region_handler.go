package api

import (
	"github.com/gorilla/mux"
	"github.com/aimir/goplaces/repository"
	"net/http"
	"strconv"
	"encoding/json"
)

func GetCountryRegionsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, _ := strconv.Atoi(params["country"])
	json.NewEncoder(w).Encode(repository.GetAllRegionsByCountryID(country))
}

func GetCountryCodeRegionsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(repository.GetAllRegionsByCountryCode(params["country"]))
}

func GetRegionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(repository.GetRegion(id))
}
