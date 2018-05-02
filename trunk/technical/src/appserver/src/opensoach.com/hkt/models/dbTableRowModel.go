package models

import (
	"time"
)

type DBSplHktChartTasksTableRowModel struct {
	ChartId   int64     `db:"chart_id_fk" dbattr:"pri"  json:"chartid"`
	TaskId    int64     `db:"task_id_fk" dbattr:"pri"  json:"taskid"`
	TaskOrder *int      `db:"task_order" json:"taskorder"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHktChartTableRowModel struct {
	ChartId     int64     `db:"id" dbattr:"pri,auto"  json:"chartid"`
	CpmId       int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId       int64     `db:"spc_id_fk" json:"spcid"`
	ChartName   string    `db:"chart_name" json:"chartname"`
	ChartType   int       `db:"chart_type" json:"charttype"`
	ChartConfig int64     `db:"chart_config" json:"chartconfig"`
	CreatedOn   time.Time `db:"created_on" json:"createdon"`
	UpdatedOn   time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHktChartTxnTableRowModel struct {
	ChartTxnId int64     `db:"id" dbattr:"pri,auto"  json:"charttxnid"`
	ChartId    int64     `db:"chart_id_fk" json:"chartid"`
	TaskId     int64     `db:"task_id_fk" json:"taskid"`
	Slot       int       `db:"slot" json:"slot"`
	TaskState  int       `db:"task_state" json:"taskstate"`
	EntryTime  time.Time `db:"entry_time" json:"entrytime"`
	TaskTxnDay time.Time `db:"task_txn_day" json:"tasktxnday"`
	CreatedOn  time.Time `db:"created_on" json:"createdon"`
	UpdatedOn  time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHktDevSpMappingRowModel struct {
	DevId     int64     `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	SpId      int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHktFieldOperatorTableRowModel struct {
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

type DBSplHktFopSpTableRowModel struct {
	FopId     int64     `db:"fop_id_fk" dbattr:"pri"  json:"fopid"`
	SpId      int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHktSpChartsTableRowModel struct {
	ChartId   int64     `db:"chart_id_fk" dbattr:"pri"  json:"chartid"`
	SpId      int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHktSpComplaintTableRowModel struct {
	ComplaintId    int64      `db:"id" dbattr:"pri,auto"  json:"Complaintid"`
	SpId           int64      `db:"sp_id_fk" json:"spid"`
	ComplaintTitle string     `db:"complaint_title" json:"complainttitle"`
	Description    *string    `db:"description" json:"description"`
	ComplaintBy    string     `db:"complaint_by" json:"complaintby"`
	MobileNo       *string    `db:"mobile_no" json:"mobileno"`
	EmailId        *string    `db:"email_id" json:"emailid"`
	EmployeeId     *string    `db:"employee_id" json:"employeeid"`
	RaisedOn       time.Time  `db:"raised_on" json:"raisedon"`
	ComplaintState int        `db:"complaint_state" json:"complaintstate"`
	ClosedOn       *time.Time `db:"closed_on" json:"closedon"`
	Remarks        *string    `db:"remarks" json:"remarks"`
	CreatedOn      time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn      time.Time  `db:"updated_on" json:"updatedon"`
}

type DBSplHktTaskLibTableRowModel struct {
	TaskLibId int64     `db:"id" dbattr:"pri,auto"  json:"tasklibid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId     *int64    `db:"spc_id_fk" json:"spcid"`
	TaskName  string    `db:"task_name" json:"taskname"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}
