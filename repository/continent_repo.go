package repository

import (
	"github.com/aimir/goplaces/database"
	"github.com/aimir/goplaces/model"
)

//GetAllContinents gets all containers
func GetAllContinents() []*model.Continent {
	var cts []*model.Continent
	database.GetDB().All(&cts)

	return cts
}

//GetContinent gets container by ID
func GetContinent(id int) model.Continent {
	var ct model.Continent
	database.GetDB().One("ID", id, &ct)

	return ct
}

//GetContinentByCode gets container by code
func GetContinentByCode(code string) model.Continent {
	var ct model.Continent
	database.GetDB().One("Code", code, &ct)

	return ct
}
