package repository

import (
	"fmt"
	"github.com/aimir/goplaces/database"
	"github.com/aimir/goplaces/model"
)

//GetAllCities gets all cities
func GetAllCities() []*model.City {
	var cs []*model.City
	database.GetDB().All(&cs)

	return cs
}

//GetAllCitiesByCountryID gets cities by country ID
func GetAllCitiesByCountryID(countryId int) []*model.City {
	var cs []*model.City
	database.GetDB().Find("CountryID", countryId, &cs)

	return cs
}

//GetAllCitiesByCountryCode gets cities by country code
func GetAllCitiesByCountryCode(countryCode string) []*model.City {
	var cs []*model.City
	database.GetDB().Find("CountryCode", countryCode, &cs)

	return cs
}

//GetAllCitiesByCountryIDByRegionID gets cities by country ID and region ID
func GetAllCitiesByCountryIDByRegionID(countryId int, regionId int) []*model.City {
	IDKey := fmt.Sprintf("%d_%d", countryId, regionId)
	var cs []*model.City
	database.GetDB().Find("IDKey", IDKey, &cs)

	return cs
}

//GetAllCitiesByCountryCodeByRegionCode gets cities by country code and region code
func GetAllCitiesByCountryCodeByRegionCode(countryCode string, regionCode string) []*model.City {
	CodeKey := fmt.Sprintf("%s_%s", countryCode, regionCode)
	var cs []*model.City
	database.GetDB().Find("CodeKey", CodeKey, &cs)

	return cs
}

//GetCity gets city by ID
func GetCity(id int) model.City {
	var c model.City
	database.GetDB().One("ID", id, &c)

	return c
}

//GetCityView gets city viewer by city ID
func GetCityView(id int) model.CityView {
	city := GetCity(id)
	country := GetCountry(city.CountryID)
	region := GetRegion(city.RegionID)
	return model.CityView{
		ID:          city.ID,
		CountryID:   country.ID,
		CountryCode: country.Code,
		CountryName: country.Name,
		RegionID:    region.ID,
		RegionCode:  region.Code,
		RegionName:  region.Name,
		Name:        city.Name,
	}
}
