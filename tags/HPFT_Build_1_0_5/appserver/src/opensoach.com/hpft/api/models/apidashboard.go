package models

import (
	"time"
)

type APIDashboardDeviceSummaryResponse struct {
	TotalDevices   int `json:"total"`
	Onlinedevices  int `json:"online"`
	Offlinedevices int `json:"offline"`
}

type APIDashboardLocationSummaryResponse struct {
	Total  int `json:"total"`
	Active int `json:"active"`
	InUse  int `json:"inuse"`
}

type APIDashboardFeedbackFilterModel struct { //Parameter mapping is done from client side
	CPMID     int64      `db:"cpm_id_fk"`
	SPId      *int       `db:"sp_id_fk"  json:"spid"`
	StartTime *time.Time `db:"raised_on"  json:"startdate"`
	EndTime   *time.Time `db:"raised_on"  json:"enddate"`
}

type APIDashboardComplaintFilterModel struct { //Parameter mapping is done from client side
	CPMID     int64      `db:"cpm_id_fk"`
	SPId      *int       `db:"sp_id_fk"  json:"spid"`
	StartTime *time.Time `db:"raised_on"  json:"startdate"`
	EndTime   *time.Time `db:"raised_on"  json:"enddate"`
}

type APIDashboardFeedbackResponse struct {
	Rating1 int `json:"rating1"`
	Rating2 int `json:"rating2"`
	Rating3 int `json:"rating3"`
	Rating4 int `json:"rating4"`
	Rating5 int `json:"rating5"`
}

type APIDashboardTaskRequest struct {
	SPId      *int64     `db:"sp_id_fk"  json:"spid"`
	StartTime *time.Time `db:"raised_on"  json:"startdate"`
	EndTime   *time.Time `db:"raised_on"  json:"enddate"`
}

type APIDashboardTaskResponse struct {
	Ontime  int `json:"ontime"`
	Delayed int `json:"delayed"`
}

type APIDashboardComplaintResponse struct {
	Open       int `json:"open"`
	Close      int `json:"closed"`
	Inprogress int `json:"inprogress"`
}

type APIFeedbacksPerMonthRequest struct {
	SpID      *int64     `json:"spid"`
	StartDate *time.Time `json:"startdate"`
	EndDate   *time.Time `json:"enddate"`
}

type APIComplaintsByMonthRequest struct {
	SpID      *int64     `json:"spid"`
	StartDate *time.Time `json:"startdate"`
	EndDate   *time.Time `json:"enddate"`
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

type APITaskByMonthRequest struct {
	SpID      *int64     `json:"spid"`
	StartDate *time.Time `json:"startdate"`
	EndDate   *time.Time `json:"enddate"`
}

type APITopFeedbacksResponse struct {
	FeedbackId      int64   `db:"id" dbattr:"pri,auto"  json:"feedbackid"`
	Feedback        int     `db:"feedback" json:"feedback"`
	FeedbackComment *string `db:"feedback_comment" json:"feedbackcomment"`
}

type APITopFeedbacksRequest struct {
	SpID          *int64 `json:"spid"`
	NoOfFeedbacks int    `json:"nooffeedbacks"`
}

type APIDashboardPatientFilterModel struct { //Parameter mapping is done from client side
	CpmId int64 `db:"patient.cpm_id_fk"`
	SPId  *int  `db:"sp_id_fk"  json:"spid"`
}

type APIDashboardPatientResponse struct {
	Admitted   int `json:"admitted"`
	Discharged int `json:"discharged"`
}

type APIPatientHospitalisedByMonthRequest struct {
	SpID      *int64     `json:"spid"`
	StartDate *time.Time `json:"startdate"`
	EndDate   *time.Time `json:"enddate"`
}
