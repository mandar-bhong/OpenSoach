package models

type DBDashBoardDeviceSummaryDataModel struct {
	ConnectionState int `db:"connection_state"  json:"connstate"`
	Count           int `db:"count" json:"count"`
}
