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

//GetContinentsHandler handle request to get all continents
func GetContinentsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repository.GetAllContinents())
}

//GetContinentHandler handle request to get continent by ID
func GetContinentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(repository.GetContinent(id))
}

//GetContinentByCodeHandler handle request to get continent by code
func GetContinentByCodeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(repository.GetContinentByCode(params["code"]))
}

func sortContinents(c []*model.Continent, r *http.Request) {
	queryValues := r.URL.Query()
	sortBy := queryValues.Get("sort")
	if sortBy == "" {
		sortBy = "name"
	}
	desc := strings.ToLower(queryValues.Get("direct")) == "desc"
	switch strings.ToLower(sortBy) {
	case "code":
		model.SortContinentsByCode(c, desc)
	default:
		model.SortContinentsByName(c, desc)
	}
}
