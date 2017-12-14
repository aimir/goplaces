package repository

import (
	"github.com/aimir/goplaces/model"
	"github.com/aimir/goplaces/database"
)

func GetAllCountries() []model.Country {
	var cs []model.Country
	database.GetDB().All(&cs)

	return cs
}

func GetAllCountriesByContinentID(continentId int) []model.Country {
	var cs []model.Country
	database.GetDB().Find("ContinentID", continentId, &cs)

	return cs
}

func GetAllCountriesByContinentCode(continentCode string) []model.Country {
	var cs []model.Country
	database.GetDB().Find("ContinentCode", continentCode, &cs)

	return cs
}

func GetCountry(id int) model.Country {
	var c model.Country
	database.GetDB().One("ID", id, &c)

	return c
}

func GetCountryByCode(code string) model.Country {
	var c model.Country
	database.GetDB().One("Code", code, &c)

	return c
}
