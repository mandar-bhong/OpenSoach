package models

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
