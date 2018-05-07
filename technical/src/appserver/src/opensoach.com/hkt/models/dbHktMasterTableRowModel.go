package models

import "time"

type DBSplHktMasterSpcTaskLibTableRowModel struct {
	SpcId     int64     `db:"spc_id_fk" dbattr:"pri"  json:"spcid"`
	MtaskId   int64     `db:"mtask_id_fk" dbattr:"pri"  json:"mtaskid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHktMasterTaskLibTableRowModel struct {
	TaskId    int64     `db:"id" dbattr:"pri,auto"  json:"taskid"`
	TaskName  string    `db:"task_name" json:"taskname"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplProdMasterConfigRowModel struct {
	ConfigKey   string    `db:"config_key" dbattr:"pri"  json:"configkey"`
	ConfigValue string    `db:"config_value" json:"configvalue"`
	CreatedOn   time.Time `db:"created_on" json:"createdon"`
	UpdatedOn   time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplProdMasterServConfTypeTableRowModel struct {
	ServConfId   int64     `db:"id" dbattr:"pri,auto"  json:"servconfid"`
	ConfTypeCode string    `db:"conf_type_code" json:"conftypecode"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplProdMasterSpCategoryTableRowModel struct {
	SpcId     int64     `db:"id" dbattr:"pri,auto"  json:"spcid"`
	SpcName   string    `db:"spc_name" json:"spcname"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}
