package models

import (
	"time"

	hktmodels "opensoach.com/hkt/models"
)

type APIViewReportRequestModel struct {
	ReportRequest []hktmodels.DBReportRequestDataModel `json:"reportreq"`
}

type APIGenerateReportRequestModel struct {
	APIViewReportRequestModel
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}
