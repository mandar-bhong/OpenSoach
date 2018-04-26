package models

type APITaskAddRequest struct {
	Name        string `json:"name"`
	Discription string `json:"desc"`
}
