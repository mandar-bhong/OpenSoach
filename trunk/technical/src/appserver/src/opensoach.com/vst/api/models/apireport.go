package models

import (
	hktmodels "opensoach.com/vst/models"
)

type APIViewReportRequestModel struct {
	ReportRequest []hktmodels.DBReportRequestDataModel `json:"reportreq"`
}

type APIGenerateReportRequestModel struct {
	APIViewReportRequestModel
	ReportFileFormat string
}
