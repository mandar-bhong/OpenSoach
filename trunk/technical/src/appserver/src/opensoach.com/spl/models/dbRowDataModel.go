package models

type DBTotalRecordsModel struct {
	TotalRecords int `db:"count" json:"count"`
}
