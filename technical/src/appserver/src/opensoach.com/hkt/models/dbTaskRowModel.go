package models

type DBTaskLibRowModel struct {
	CPMID        int64  `db:"cpm_id_fk" json:"cpmid"`
	SPCategoryID int64  `db:"spc_id_fk" json:"spcid"`
	TaskName     string `db:"task_name" json:"taskname"`
	Description  string `db:"short_desc" json:"desc"`
}
