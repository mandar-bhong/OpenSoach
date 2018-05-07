package models

type DBTaskLibDataModel struct {
	SpcId     *int64  `db:"spc_id_fk" json:"spcid"`
	TaskName  string  `db:"task_name" json:"taskname"`
	ShortDesc *string `db:"short_desc" json:"shortdesc"`
}

type DBTaskLibInsertRowModel struct {
	CPMID int64 `db:"cpm_id_fk" json:"cpmid"`
	DBTaskLibDataModel
}

type DBTaskLibUpdateRowModel struct {
	TaskLibId int64   `db:"id" dbattr:"pri,auto"  json:"tasklibid"`
	CpmId     int64   `db:"cpm_id_fk" json:"cpmid"`
	SpcId     *int64  `db:"spc_id_fk" json:"spcid"`
	TaskName  string  `db:"task_name" json:"taskname"`
	ShortDesc *string `db:"short_desc" json:"shortdesc"`
}
