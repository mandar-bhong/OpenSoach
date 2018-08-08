package models

import "time"

type DBSplHktSpComplaintTableRowModel struct {
	ComplaintId    int64      `db:"id" dbattr:"pri,auto"  json:"complaintid"`
	CpmId          int64      `db:"cpm_id_fk" json:"cpmid"`
	SpId           int64      `db:"sp_id_fk" json:"spid"`
	ComplaintTitle string     `db:"complaint_title" json:"complainttitle"`
	Description    *string    `db:"description" json:"description"`
	ComplaintBy    string     `db:"complaint_by" json:"complaintby"`
	MobileNo       *string    `db:"mobile_no" json:"mobileno"`
	EmailId        *string    `db:"email_id" json:"emailid"`
	EmployeeId     *string    `db:"employee_id" json:"employeeid"`
	Severity       *int       `db:"severity" json:"severity"`
	RaisedOn       time.Time  `db:"raised_on" json:"raisedon"`
	ComplaintState int        `db:"complaint_state" json:"complaintstate"`
	ClosedOn       *time.Time `db:"closed_on" json:"closedon"`
	Remarks        *string    `db:"remarks" json:"remarks"`
	CreatedOn      time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn      time.Time  `db:"updated_on" json:"updatedon"`
}

type DBSplHktTaskLibTableRowModel struct {
	TaskId    int64     `db:"id" dbattr:"pri,auto"  json:"taskid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId     int64     `db:"spc_id_fk" json:"spcid"`
	TaskName  string    `db:"task_name" json:"taskname"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeCpmTableRowModel struct {
	CpmId int64 `db:"cpm_id_fk" dbattr:"pri,auto"  json:"cpmid"`
}

type DBSplNodeDevSpMappingRowModel struct {
	DevId     int64     `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	SpId      int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeDevTableRowModel struct {
	DevId     int64     `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	DevName   string    `db:"dev_name" json:"devname"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	Serialno  string    `db:"serialno" json:"serialno"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeDevStatusTableRowModel struct {
	DevId                int64     `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	ConnectionState      int       `db:"connection_state" json:"connectionstate"`
	ConnectionStateSince time.Time `db:"connection_state_since" json:"connectionstatesince"`
	SyncState            int       `db:"sync_state" json:"syncstate"`
	SyncStateSince       time.Time `db:"sync_state_since" json:"syncstatesince"`
	BatteryLevel         int       `db:"battery_level" json:"batterylevel"`
	BatteryLevelSince    time.Time `db:"battery_level_since" json:"batterylevelsince"`
	CreatedOn            time.Time `db:"created_on" json:"createdon"`
	UpdatedOn            time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeFieldOperatorTableRowModel struct {
	FopId     int64     `db:"id" dbattr:"pri,auto"  json:"fopid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	Fopcode   string    `db:"fopcode" json:"fopcode"`
	FopName   *string   `db:"fop_name" json:"fopname"`
	MobileNo  string    `db:"mobile_no" json:"mobileno"`
	EmailId   *string   `db:"email_id" json:"emailid"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	FopState  int       `db:"fop_state" json:"fopstate"`
	FopArea   int       `db:"fop_area" json:"foparea"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeFopSpTableRowModel struct {
	FopId     int64     `db:"fop_id_fk" dbattr:"pri"  json:"fopid"`
	SpId      int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeServiceConfTableRowModel struct {
	ServConfId   int64     `db:"id" dbattr:"pri,auto"  json:"servconfid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	ConfTypeCode string    `db:"conf_type_code" json:"conftypecode"`
	ServConfName string    `db:"serv_conf_name" json:"servconfname"`
	ShortDesc    *string   `db:"short_desc" json:"shortdesc"`
	ServConf     string    `db:"serv_conf" json:"servconf"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeServiceInTxnTableRowModel struct {
	ServInTxnId int64     `db:"id" dbattr:"pri,auto"  json:"servintxnid"`
	CpmId       int64     `db:"cpm_id_fk" json:"cpmid"`
	ServInId    int64     `db:"serv_in_id_fk" json:"servinid"`
	TxnData     string    `db:"txn_data" json:"txndata"`
	TxnDate     time.Time `db:"txn_date" json:"txndate"`
	CreatedOn   time.Time `db:"created_on" json:"createdon"`
	UpdatedOn   time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeServiceInstanceTableRowModel struct {
	ServInId   int64     `db:"id" dbattr:"pri,auto"  json:"servinid"`
	CpmId      int64     `db:"cpm_id_fk" json:"cpmid"`
	ServConfId int64     `db:"serv_conf_id_fk" json:"servconfid"`
	SpId       int64     `db:"sp_id_fk" json:"spid"`
	CreatedOn  time.Time `db:"created_on" json:"createdon"`
	UpdatedOn  time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeSpCategoryTableRowModel struct {
	SpcId     int64     `db:"id" dbattr:"pri,auto"  json:"spcid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcName   string    `db:"spc_name" json:"spcname"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeSpTableRowModel struct {
	SpId         int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	SpName       string    `db:"sp_name" json:"spname"`
	ShortDesc    *string   `db:"short_desc" json:"shortdesc"`
	SpState      int       `db:"sp_state" json:"spstate"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeFeedbackTableRowModel struct {
	FeedbackId      int64     `db:"id" dbattr:"pri,auto"  json:"feedbackid"`
	CpmId           int64     `db:"cpm_id_fk" json:"cpmid"`
	SpId            int64     `db:"sp_id_fk" json:"spid"`
	Feedback        int       `db:"feedback" json:"feedback"`
	FeedbackComment *string   `db:"feedback_comment" json:"feedbackcomment"`
	RaisedOn        time.Time `db:"raised_on" json:"raisedon"`
	CreatedOn       time.Time `db:"created_on" json:"createdon"`
}

type DBSplNodeReportTemplateTableRowModel struct {
	ReportId          int64  `db:"id" dbattr:"pri,auto"  json:"reportid"`
	ReportCode        string `db:"report_code" json:"reportcode"`
	ReportDesc        string `db:"report_desc" json:"reportdesc"`
	ReportHeader      string `db:"report_header" json:"reportheader"`
	ReportQueryParams string `db:"report_query_params" json:"reportqueryparams"`
	ReportQuery       string `db:"report_query" json:"reportquery"`
}

type DBSplHpftPatientMasterTableRowModel struct {
	PatientId             int64      `db:"id" dbattr:"pri,auto"  json:"patientid"`
	CpmIdFk               int64      `db:"cpm_id_fk" json:"cpmidfk"`
	PatientDetails        string     `db:"patient_details" json:"patientdetails"`
	MedicalDetails        string     `db:"medical_details" json:"medicaldetails"`
	PatientFileTemplateID int64      `db:"patient_file_template" json:"patientfiletemplate"`
	SpId                  int64      `db:"sp_id_fk" json:"spid"`
	ServInId              int64      `db:"serv_in_id_fk" json:"servinid"`
	Status                int        `db:"status" json:"status"`
	DischargedOn          *time.Time `db:"discharged_on" json:"dischargedon"`
}
