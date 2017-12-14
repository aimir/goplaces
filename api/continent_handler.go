package api

import (
	"github.com/aimir/goplaces/repository"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetContinentsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repository.GetAllContinents())
}

func GetContinentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(repository.GetContinent(id))
}

func GetContinentByCodeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(repository.GetContinentByCode(params["code"]))
}
