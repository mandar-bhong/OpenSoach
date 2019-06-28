package models

type APIProductGetRequest struct {
	ProductID int `json:"cpmid"`
}

type APIProductSelectRequest struct {
	ProductID int64 `json:"cpmid"`
}
