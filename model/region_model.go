package model

import "sort"

type Region struct {
	ID          int    `json:"id"`
	Code        string `storm:"index" json:"code"`
	CountryID   int    `storm:"index" json:"-"`
	CountryCode string `storm:"index" json:"-"`
	Name        string `json:"name"`
}

type regions []*Region

func (s regions) Len() int { return len(s) }
func (s regions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type regionsByName struct {regions}

func (s regionsByName) Less(i, j int) bool {
	return s.regions[i].Name < s.regions[j].Name
}

type regionsByCode struct {regions}

func (s regionsByCode) Less(i, j int) bool {
	return s.regions[i].Code < s.regions[j].Code
}

func SortRegionsByName(r []*Region, desc bool) {
	if desc {
		sort.Sort(sort.Reverse(regionsByName{r}))
	} else {
		sort.Sort(regionsByName{r})
	}
}

func SortRegionsByCode(r []*Region, desc bool) {
	if desc {
		sort.Sort(sort.Reverse(regionsByCode{r}))
	} else {
		sort.Sort(regionsByCode{r})
	}
}
