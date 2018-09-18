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
	Year    int `db:"year" json:"year"`
	Month   int `db:"month" json:"month"`
	Rating1 int `db:"rating1" json:"rating1"`
	Rating2 int `db:"rating2" json:"rating2"`
	Rating3 int `db:"rating3" json:"rating3"`
	Rating4 int `db:"rating4" json:"rating4"`
	Rating5 int `db:"rating5" json:"rating5"`
}

type DBFeedbacksPerMonthFilterDataModel struct {
	SpId  *int64 `db:"sp_id_fk" json:"spid"`
	CpmId int64  `db:"cpm_id_fk" json:"cpmid"`
}

type DBNoOfComplaintsPerMonthsFilterDataModel struct {
	SpId  *int64 `db:"sp_id_fk" json:"spid"`
	CpmId int64  `db:"cpm_id_fk" json:"cpmid"`
}

type DBNoOfComplaintsPerMonthDataModel struct {
	Year       int `db:"year" json:"year"`
	Month      int `db:"month" json:"month"`
	Open       int `db:"open" json:"open"`
	Closed     int `db:"closed" json:"closed"`
	InProgress int `db:"inprogress" json:"inprogress"`
}

type DBTopComplaintsFilterDataModel struct {
	SpId           *int64 `db:"sp_id_fk" json:"spid"`
	CpmId          int64  `db:"cpm_id_fk" json:"cpmid"`
	ComplaintState int    `db:"complaint_state" json:"complaintstate"`
}

type DBTaskPerMonthFilterDataModel struct {
	SpId  *int64 `db:"sp_id_fk" json:"spid"`
	CpmId int64  `db:"cpm_id_fk" json:"cpmid"`
}

type DBTaskSummaryPerMonthDataModel struct {
	Year    int `db:"year" json:"year"`
	Month   int `db:"month" json:"month"`
	Ontime  int `db:"ontime" json:"ontime"`
	Delayed int `db:"delay" json:"delayed"`
}

type DBTopFeedbackFilterDataModel struct {
	SpId  *int64 `db:"sp_id_fk" json:"spid"`
	CpmId int64  `db:"cpm_id_fk" json:"cpmid"`
}

type DBSnapshotDataModel struct {
	TxnDate time.Time `db:"txn_date" json:"lastactiontime"`
	Status  int       `db:"status" json:"status"`
	Count   int       `db:"count" json:"count"`
}

type DBAverageTimeDataModel struct {
	WaitTime        *float64 `db:"waittime" json:"waittime"`
	JobCreationTime *float64 `db:"jobcreationtime" json:"jobcreationtime"`
	JobExeTime      *float64 `db:"jobexetime" json:"jobexetime"`
	DeliveryTime    *float64 `db:"deliverytime" json:"deliverytime"`
}
