package repository

import (
	"github.com/aimir/goplaces/database"
	"github.com/aimir/goplaces/model"
)

//Get all regions
func GetAllRegions() []*model.Region {
	var rs []*model.Region
	database.GetDB().All(&rs)

	return rs
}

//Get regions by country ID
func GetAllRegionsByCountryID(countryId int) []*model.Region {
	var rs []*model.Region
	database.GetDB().Find("CountryID", countryId, &rs)

	return rs
}

//Get regions by country code
func GetAllRegionsByCountryCode(countryCode string) []*model.Region {
	var rs []*model.Region
	database.GetDB().Find("CountryCode", countryCode, &rs)

	return rs
}

//Get region by ID
func GetRegion(id int) model.Region {
	var r model.Region
	database.GetDB().One("ID", id, &r)

	return r
}
