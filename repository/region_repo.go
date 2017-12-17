package repository

import (
	"github.com/aimir/goplaces/database"
	"github.com/aimir/goplaces/model"
)

func GetAllRegions() []*model.Region {
	var rs []*model.Region
	database.GetDB().All(&rs)

	return rs
}

func GetAllRegionsByCountryID(countryId int) []*model.Region {
	var rs []*model.Region
	database.GetDB().Find("CountryID", countryId, &rs)

	return rs
}

func GetAllRegionsByCountryCode(countryCode string) []*model.Region {
	var rs []*model.Region
	database.GetDB().Find("CountryCode", countryCode, &rs)

	return rs
}

func GetRegion(id int) model.Region {
	var r model.Region
	database.GetDB().One("ID", id, &r)

	return r
}
