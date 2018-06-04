package models

type DBEPSPDataModel struct {
	ID           int64  `db:"sp_id_fk" json:"spid"`
	Name         string `db:"sp_name" json:"spname"`
	CategoryName string `db:"spc_name" json:"spcname"`
}

type DBEPSPServConfDataModel struct {
	ServInId     int64  `db:"id" dbattr:"pri,auto"  json:"servinid"`
	ServConfId   int64  `db:"serv_conf_id_fk" json:"servconfid"`
	ConfTypeCode string `db:"conf_type_code" json:"conftypecode"`
	ServConfName string `db:"serv_conf_name" json:"servconfname"`
	ServConf     string `db:"serv_conf" json:"servconf"`
}

type DBEPSPFieldOperatorDataModel struct {
	FopId    int64   `db:"fop_id_fk" json:"fopid"`
	Fopcode  string  `db:"fopcode" json:"fopcode"`
	FopName  *string `db:"fop_name" json:"fopname"`
	FopState int     `db:"fop_state" json:"fopstate"`
	FopArea  int     `db:"fop_area" json:"foparea"`
}
