package models

type APICustSpCountUpdateRequest struct {
	CpmId       int64 `json:cpmid`
	UpdateCount int   `json:"updatecount"`
}

type APICustCpmListRequest struct {
	ProdCode string `json:"prodcode"`
}
