package models

import (
	"time"

	hktmodels "opensoach.com/hkt/models"
)

type APIComplaintAddRequest struct {
	hktmodels.DBComplaintDataModel
}

type APITopActiveComplaintsRequest struct {
	SpID           *int64 `json:"spid"`
	NoOfComplaints int    `json:"noofcomplaints"`
}

type APITopActiveComplaintsResponse struct {
	ComplaintId    int64     `db:"id" dbattr:"pri,auto"  json:"complaintid"`
	ComplaintTitle string    `db:"complaint_title" json:"complainttitle"`
	RaisedOn       time.Time `db:"raised_on" json:"raisedon"`
	ComplaintState int       `db:"complaint_state" json:"complaintstate"`
	Severity       *int      `db:"severity" json:"severity"`
}

type APIComplaintsByMonthRequest struct {
	SpID      *int64     `json:"spid"`
	StartDate *time.Time `json:"startdate"`
	EndDate   *time.Time `json:"enddate"`
}
