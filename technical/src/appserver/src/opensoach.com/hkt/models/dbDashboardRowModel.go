package models

import "time"

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

type DBDashBoardTaskDataModel struct {
	Status int `db:"status" json:"status"`
	Count  int `db:"count" json:"count"`
}

type DBTaskSummaryFilterDataModel struct {
	SpId  *int64 `db:"sp_id_fk" json:"spid"`
	CpmId int64  `db:"serv_in_txn.cpm_id_fk" json:"cpmid"`
}

type DBDashBoardComplaintDataModel struct {
	ComplaintState int `db:"complaint_state" json:"complaintstate"`
	Count          int `db:"count" json:"count"`
}

type DBFeedbackDataModel struct {
	FeedbackId      int64     `db:"id" dbattr:"pri,auto"  json:"FeedbackId"`
	SpId            int64     `db:"sp_id_fk" json:"spid"`
	Feedback        int       `db:"feedback" json:"feedback"`
	FeedbackComment *string   `db:"feedback_comment" json:"feedbackcomment"`
	RaisedOn        time.Time `db:"raised_on" json:"raisedon"`
	CreatedOn       time.Time `db:"created_on" json:"createdon"`
}

type DBDashBoardInUseLocationDataModel struct {
	Count int `db:"count" json:"count"`
}

type DBFeedbacksPerMonthDataModel struct {
	Year       int `db:"year" json:"year"`
	Month      int `db:"month" json:"month"`
	Rating1 int `json:"rating1"`
	Rating2 int `json:"rating2"`
	Rating3 int `json:"rating3"`
	Rating4 int `json:"rating4"`
	Rating5 int `json:"rating5"`
}

type DBFeedbacksPerMonthFilterDataModel struct {
	SpId  *int64 `db:"sp_id_fk" json:"spid"`
	CpmId int64  `db:"cpm_id_fk" json:"cpmid"`
}
