package models

type APITaskAddRequest struct {
	SPCategoryID int64  `json:"spcid"`
	Name         string `json:"name"`
	Discription  string `json:"desc"`
}
