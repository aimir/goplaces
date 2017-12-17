package model

import "sort"

//Country data type
type Country struct {
	ID            int    `json:"id"`
	Code          string `storm:"unique" json:"code"`
	ContinentID   int    `storm:"index" json:"-"`
	ContinentCode string `storm:"index" json:"-"`
	Name          string `storm:"unique" json:"name"`
}

type countries []*Country

func (s countries) Len() int      { return len(s) }
func (s countries) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type countryByName struct{ countries }

func (s countryByName) Less(i, j int) bool {
	return s.countries[i].Name < s.countries[j].Name
}

type countryByCode struct{ countries }

func (s countryByCode) Less(i, j int) bool {
	return s.countries[i].Code < s.countries[j].Code
}

//Sort countries slice by name
func SortCountriesByName(c []*Country, desc bool) {
	if desc {
		sort.Sort(sort.Reverse(countryByName{c}))
	} else {
		sort.Sort(countryByName{c})
	}
}

//Sort countries slicer by code
func SortCountriesByCode(c []*Country, desc bool) {
	if desc {
		sort.Sort(sort.Reverse(countryByCode{c}))
	} else {
		sort.Sort(countryByCode{c})
	}
}
