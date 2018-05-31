package models

type DBGetReportDataModel struct {
	ReportId     int64      `db:"id" dbattr:"pri,auto"  json:"reportid"`
	ReportCode   string     `db:"report_code" json:"reportcode"`
	ReportDesc   string     `db:"report_desc" json:"reportdesc"`
	ReportHeader string     `db:"report_header" json:"reportheader"`
	ReportData   [][]string `json:"reportdata"`
}

type DBGenerateReportRequestDataModel struct {
	ReportID int64  `db:"id" json:"reportid"`
	Language string `json:"lang"`
}

type DBReportTemplateShortDataModel struct {
	ReportId   int64  `db:"id" dbattr:"pri,auto"  json:"reportid"`
	ReportCode string `db:"report_code" json:"reportcode"`
	ReportDesc string `db:"report_desc" json:"reportdesc"`
}

type ReportHeaderModel struct {
	En []string `json:"en"`
	Hi []string `json:"hi"`
}
