package repository

import (
	"fmt"
	"github.com/aimir/goplaces/database"
	"github.com/aimir/goplaces/model"
)

func GetAllCities() []*model.City {
	var cs []*model.City
	database.GetDB().All(&cs)

	return cs
}

func GetAllCitiesByCountryID(countryId int) []*model.City {
	var cs []*model.City
	database.GetDB().Find("CountryID", countryId, &cs)

	return cs
}

func GetAllCitiesByCountryCode(countryCode string) []*model.City {
	var cs []*model.City
	database.GetDB().Find("CountryCode", countryCode, &cs)

	return cs
}

func GetAllCitiesByCountryIDByRegionID(countryId int, regionId int) []*model.City {
	IDKey := fmt.Sprintf("%d_%d", countryId, regionId)
	var cs []*model.City
	database.GetDB().Find("IDKey", IDKey, &cs)

	return cs
}

func GetAllCitiesByCountryCodeByRegionCode(countryCode string, regionCode string) []*model.City {
	CodeKey := fmt.Sprintf("%s_%s", countryCode, regionCode)
	var cs []*model.City
	database.GetDB().Find("CodeKey", CodeKey, &cs)

	return cs
}

func GetCity(id int) model.City {
	var c model.City
	database.GetDB().One("ID", id, &c)

	return c
}

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
