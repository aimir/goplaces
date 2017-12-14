package model

type Continent struct {
	ID   int    `json:"id"`
	Code string `storm:"unique" json:"code"`
	Name string `storm:"unique" json:"name"`
}
