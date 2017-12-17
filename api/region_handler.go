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

func GetCountryRegionsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, _ := strconv.Atoi(params["country"])
	regions := repository.GetAllRegionsByCountryID(country)
	sortRegions(regions, r)
	json.NewEncoder(w).Encode(regions)
}

func GetCountryCodeRegionsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regions := repository.GetAllRegionsByCountryCode(params["country"])
	sortRegions(regions, r)
	json.NewEncoder(w).Encode(regions)
}

func GetRegionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(repository.GetRegion(id))
}

func sortRegions(rg []*model.Region, r *http.Request) {
	queryValues := r.URL.Query()
	sortBy := queryValues.Get("sort")
	if sortBy == "" {
		sortBy = "name"
	}
	desc := strings.ToLower(queryValues.Get("direct")) == "desc"
	switch strings.ToLower(sortBy) {
	case "code":
		model.SortRegionsByCode(rg, desc)
	default:
		model.SortRegionsByName(rg, desc)
	}
}
