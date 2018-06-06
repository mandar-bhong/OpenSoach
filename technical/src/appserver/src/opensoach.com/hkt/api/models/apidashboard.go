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

type APIDashboardComplaintResponse struct {
	Open       int `json:"open"`
	Close      int `json:"closed"`
	Inprogress int `json:"inprogress"`
}
