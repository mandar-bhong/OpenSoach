package models

type DBDashBoardDeviceSummaryDataModel struct {
	ConnectionState int `db:"connection_state"  json:"connstate"`
	Count           int `db:"count" json:"count"`
}

type DBDashBoardLocationSummaryDataModel struct {
	State int `db:"sp_state"  json:"spstate"`
	Count int `db:"count" json:"count"`
}
