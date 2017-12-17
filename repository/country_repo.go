package repository

import (
	"github.com/aimir/goplaces/database"
	"github.com/aimir/goplaces/model"
)

//Get all countries
func GetAllCountries() []*model.Country {
	var cs []*model.Country
	database.GetDB().All(&cs)

	return cs
}

//Get countries by continent ID
func GetAllCountriesByContinentID(continentId int) []*model.Country {
	var cs []*model.Country
	database.GetDB().Find("ContinentID", continentId, &cs)

	return cs
}

//Get countries by continent code
func GetAllCountriesByContinentCode(continentCode string) []*model.Country {
	var cs []*model.Country
	database.GetDB().Find("ContinentCode", continentCode, &cs)

	return cs
}

//Get country data by ID
func GetCountry(id int) model.Country {
	var c model.Country
	database.GetDB().One("ID", id, &c)

	return c
}

//Get country data by code
func GetCountryByCode(code string) model.Country {
	var c model.Country
	database.GetDB().One("Code", code, &c)

	return c
}
