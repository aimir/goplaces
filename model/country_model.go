package model

type Country struct {
	ID            int    `json:"id"`
	Code          string `storm:"unique" json:"code"`
	ContinentID   int    `storm:"index" json:"-"`
	ContinentCode string `storm:"index" json:"-"`
	Name          string `storm:"unique" json:"name"`
}
