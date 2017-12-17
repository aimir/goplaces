package model

import "sort"

//Continent data structure
type Continent struct {
	ID   int    `json:"id"`
	Code string `storm:"unique" json:"code"`
	Name string `storm:"unique" json:"name"`
}

type continents []*Continent

func (s continents) Len() int      { return len(s) }
func (s continents) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type continentsByName struct{ continents }

func (s continentsByName) Less(i, j int) bool {
	return s.continents[i].Name < s.continents[j].Name
}

type continentsByCode struct{ continents }

func (s continentsByCode) Less(i, j int) bool {
	return s.continents[i].Code < s.continents[j].Code
}

//SortContinentsByName sorts containers by name
func SortContinentsByName(c []*Continent, desc bool) {
	if desc {
		sort.Sort(sort.Reverse(continentsByName{c}))
	} else {
		sort.Sort(continentsByName{c})
	}
}

//SortContinentsByCode sorts containers by code
func SortContinentsByCode(c []*Continent, desc bool) {
	if desc {
		sort.Sort(sort.Reverse(continentsByCode{c}))
	} else {
		sort.Sort(continentsByCode{c})
	}
}
