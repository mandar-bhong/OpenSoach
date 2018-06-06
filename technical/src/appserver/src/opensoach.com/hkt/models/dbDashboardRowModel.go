package models

type DBDashBoardDeviceSummaryDataModel struct {
	ConnectionState int `db:"connection_state"  json:"connstate"`
	Count           int `db:"count" json:"count"`
}

type DBDashBoardLocationSummaryDataModel struct {
	State int `db:"sp_state"  json:"spstate"`
	Count int `db:"count" json:"count"`
}

type DBDashBoardFeedbackRequestDataModel struct {
	CPMID *int64 `db:"cpm_id_fk"`
	SPId  *int   `db:"sp_id_fk"  json:"spid"`
}

type DBDashBoardFeedbackDataModel struct {
	Feedback int `db:"feedback" json:"feedback"`
	Count    int `db:"count" json:"count"`
}
