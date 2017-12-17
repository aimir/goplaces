package model

import "sort"

type City struct {
	ID          int    `json:"id"`
	CountryID   int    `storm:"index" json:"-"`
	CountryCode string `storm:"index" json:"-"`
	RegionID    int    `json:"-"`
	RegionCode  string `json:"-"`
	IDKey       string `storm:"index" json:"-"`
	CodeKey     string `storm:"index" json:"-"`
	Name        string `json:"name"`
}

type CityView struct {
	ID          int    `json:"id"`
	CountryID   int    `json:"countryId"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	RegionID    int    `json:"regionId"`
	RegionCode  string `json:"regionCode"`
	RegionName  string `json:"regionName"`
	Name        string `json:"name"`
}

type cities []*City

func (s cities) Len() int      { return len(s) }
func (s cities) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type citiesByName struct{ cities }

func (s citiesByName) Less(i, j int) bool {
	return s.cities[i].Name < s.cities[j].Name
}

func SortCitiesByName(c []*City, desc bool) {
	if desc {
		sort.Sort(sort.Reverse(citiesByName{c}))
	} else {
		sort.Sort(citiesByName{c})
	}
}
