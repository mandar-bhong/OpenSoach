package models

type DBEPSPDataModel struct {
	ID           int64  `db:"sp_id_fk" json:"spid"`
	Name         string `db:"sp_name" json:"name"`
	CategoryName string `db:"spc_name" json:"catname"`
}
