package repository

import (
	"github.com/aimir/goplaces/model"
	"github.com/aimir/goplaces/database"
)

func GetAllContinents() []*model.Continent {
	var cts []*model.Continent
	database.GetDB().All(&cts)

	return cts
}

func GetContinent(id int) model.Continent {
	var ct model.Continent
	database.GetDB().One("ID", id, &ct)

	return ct
}

func GetContinentByCode(code string) model.Continent {
	var ct model.Continent
	database.GetDB().One("Code", code, &ct)

	return ct
}