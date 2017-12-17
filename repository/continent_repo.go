package repository

import (
	"github.com/aimir/goplaces/database"
	"github.com/aimir/goplaces/model"
)

//Get all containers
func GetAllContinents() []*model.Continent {
	var cts []*model.Continent
	database.GetDB().All(&cts)

	return cts
}

//Get container by ID
func GetContinent(id int) model.Continent {
	var ct model.Continent
	database.GetDB().One("ID", id, &ct)

	return ct
}

//Get container by code
func GetContinentByCode(code string) model.Continent {
	var ct model.Continent
	database.GetDB().One("Code", code, &ct)

	return ct
}
