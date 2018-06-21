package models

type DBGetReportDataModel struct {
	ReportId     int64      `db:"id" dbattr:"pri,auto"  json:"reportid"`
	ReportCode   string     `db:"report_code" json:"reportcode"`
	ReportDesc   string     `db:"report_desc" json:"reportdesc"`
	ReportHeader string     `db:"report_header" json:"reportheader"`
	ReportData   [][]string `json:"reportdata"`
}

type DBGenerateReportRequestDataModel struct {
	ReportID    int64         `db:"id" json:"reportid"`
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

type DBReportLocationSummaryFilterDataModel struct {
	SpId  *int64 `db:"sp_id_fk" json:"spid"`
	CpmId int64  `db:"serv_in_txn.cpm_id_fk" json:"cpmid"`
}

type DBReportLocationSummaryDataModel struct {
	TaskName string `db:"taskname" json:"taskname"`
	Status   int    `db:"status" json:"status"`
	Count    int    `db:"count" json:"count"`
}

type ReportTaskSummary struct {
	Ontime  int `json:"ontime"`
	Delayed int `json:"delayed"`
}

type ReportTaskSummaryModel struct {
	Taskname string `json:"taskname"`
	Ontime   int    `json:"ontime"`
	Delayed  int    `json:"delayed"`
	Total    int    `json:"total"`
}
