package models

type DBGetReportDataModel struct {
	ReportID          int    `db:"id" json:"reportid"`
	ReportCode        string `db:"report_code" json:"code"`
	ReportDescription string `db:"report_desc" json:"desc"`
	ReportHeader      string `db:"report_header" json:"header"`
	ReportQuery       string `db:"report_query" json:"query"`
	ReportQueryParams string `db:"report_param" json:"params"`
}

type DBGetReportDataModel struct {
	ReportID          int        `db:"id" json:"reportid"`
	ReportData        [][]string `json:"data"`
	Header            string     `json:"header"`
	ReportCode        string     `db:"code" json:"code"`
	ReportDescription string     `db:"desc" json:"desc"`
}

type DBGenerateReportRequestDataModel struct {
	ReportID int    `db:"id" json:"reportid"`
	Language string `json:"lang"`
}
