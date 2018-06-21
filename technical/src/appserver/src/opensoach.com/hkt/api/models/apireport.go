package models

import (
	"time"

	hktmodels "opensoach.com/hkt/models"
)

type APIReportLocationSummaryRequest struct {
	ReportID  int64      `json:"reportid"`
	SpID      *int64     `json:"spid"`
	StartDate *time.Time `json:"startdate"`
	EndDate   *time.Time `json:"enddate"`
}

type APIReportLocationSummaryResponse struct {
	ReportHeader      string                             `json:"reportheader"`
	ReportTaskSummary []hktmodels.ReportTaskSummaryModel `json:"reportdata"`
}
