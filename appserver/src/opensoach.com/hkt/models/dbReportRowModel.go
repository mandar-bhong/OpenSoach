package models

type DBGetReportDataModel struct {
	ReportCode   string     `db:"report_code" json:"reportcode"`
	ReportHeader []string   `db:"report_header" json:"reportheader"`
	ReportData   [][]string `json:"reportdata"`
}

type DBReportRequestDataModel struct {
	ReportCode  string        `db:"report_code" json:"reportcode"`
	Language    string        `json:"lang"`
	QueryParams []interface{} `json:"queryparams"`
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
