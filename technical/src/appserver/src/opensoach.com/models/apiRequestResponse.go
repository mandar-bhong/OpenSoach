package models

type APIPayloadResponse struct {
	Success bool        `json:"issuccess"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type APIResponseError struct {
	Code int         `json:"code"`
	Data interface{} `json:"errordata"`
}

type APIRecordAddResponse struct {
	RecordID int64 `json:"recid"`
}

type RecordIdRequest struct {
	RecId int64 `json:"recid"`
}

type RecordIdResponse struct {
	RecId int64 `json:"recid"`
}

type DataListRequest struct {
	Limit          int         `json:"limit"`
	OrderBy        string      `json:"orderby"`
	OrderDirection string      `json:"orderdirection"`
	CurrentPage    int         `json:"page"`
	Filter         interface{} `json:"filter"`
}

type DataListResponse struct {
	// TotalRecords    int           `json:"totalrecords"`
	FilteredRecords int         `json:"filteredrecords"`
	Records         interface{} `json:"records"`
}
