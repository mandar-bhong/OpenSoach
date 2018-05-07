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
	SpcId     *int64    `db:"spc_id_fk" json:"spcid"`
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
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
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
	TxnData     int64     `db:"txn_data" json:"txndata"`
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
	SpId      int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId     int64     `db:"spc_id_fk" json:"spcid"`
	SpName    string    `db:"sp_name" json:"spname"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}
