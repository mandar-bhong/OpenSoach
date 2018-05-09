package models

type DBSpUpdateRowModel struct {
	SpId   int64  `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId  int64  `db:"cpm_id_fk" json:"cpmid"`
	SpcId  int64  `db:"spc_id_fk" json:"spcid"`
	SpName string `db:"sp_name" json:"spname"`
}

type DBSpCategoryDataModel struct {
	SpcName   string  `db:"spc_name" json:"spcname"`
	ShortDesc *string `db:"short_desc" json:"shortdesc"`
}

type DBSpCategoryInsertRowModel struct {
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
	DBSpCategoryDataModel
}

type DBFopSpDataModel struct {
	FopId int64 `db:"fop_id_fk" dbattr:"pri"  json:"fopid"`
	SpId  int64 `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
}

type DBFopSpInsertRowModel struct {
	DBFopSpDataModel
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}
