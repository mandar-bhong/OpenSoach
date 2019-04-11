package models

type APIJobStatusUpdateRequest struct {
	TokenId int64 `json:"tokenid"`
	State   int64 `json:"state"`
	Amount  int64 `json:"amount"`
}
