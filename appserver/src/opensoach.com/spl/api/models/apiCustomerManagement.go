package models

type APICustSpCountUpdateRequest struct {
	CpmId       int64 `json:cpmid`
	UpdateCount int   `json:"updatecount"`
}
