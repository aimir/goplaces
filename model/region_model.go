package model

type Region struct {
	ID          int    `json:"id"`
	Code        string `storm:"index" json:"code"`
	CountryID   int    `storm:"index" json:"-"`
	CountryCode string `storm:"index" json:"-"`
	Name        string `json:"name"`
}
