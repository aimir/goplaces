package repository

import (
	"github.com/aimir/goplaces/model"
	"github.com/aimir/goplaces/database"
	"fmt"
)

func GetAllCities() []model.City {
	var cs []model.City
	database.GetDB().All(&cs)

	return cs
}

func GetAllCitiesByCountryID(countryId int) []model.City {
	var cs []model.City
	database.GetDB().Find("CountryID", countryId, &cs)

	return cs
}

func GetAllCitiesByCountryCode(countryCode string) []model.City {
	var cs []model.City
	database.GetDB().Find("CountryCode", countryCode, &cs)

	return cs
}

func GetAllCitiesByCountryIDByRegionID(countryId int, regionId int) []model.City {
	IDKey := fmt.Sprintf("%d_%d", countryId, regionId)
	var cs []model.City
	database.GetDB().Find("IDKey", IDKey, &cs)

	return cs
}

func GetAllCitiesByCountryCodeByRegionCode(countryCode string, regionCode string) []model.City {
	CodeKey := fmt.Sprintf("%s_%s", countryCode, regionCode)
	var cs []model.City
	database.GetDB().Find("CodeKey", CodeKey, &cs)

	return cs
}

func GetCity(id int) model.City {
	var c model.City
	database.GetDB().One("ID", id, &c)

	return c
}
