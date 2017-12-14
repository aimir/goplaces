package model

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
