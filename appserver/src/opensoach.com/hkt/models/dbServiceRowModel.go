package models

type DBServiceConfDataModel struct {
	SpcId        int64   `db:"spc_id_fk" json:"spcid"`
	ConfTypeCode string  `db:"conf_type_code" json:"conftypecode"`
	ServConfName string  `db:"serv_conf_name" json:"servconfname"`
	ShortDesc    *string `db:"short_desc" json:"shortdesc"`
	ServConf     string  `db:"serv_conf" json:"servconf"`
}

type DBServiceConfInsertRowModel struct {
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
	DBServiceConfDataModel
}

type DBServiceConfUpdateRowModel struct {
	ServConfId   int64   `db:"id" dbattr:"pri,auto"  json:"servconfid"`
	CpmId        int64   `db:"cpm_id_fk" json:"cpmid"`
	ServConfName string  `db:"serv_conf_name" json:"servconfname"`
	ShortDesc    *string `db:"short_desc" json:"shortdesc"`
	ServConf     string  `db:"serv_conf" json:"servconf"`
}

type DBServiceInstanceDataModel struct {
	ServConfId int64 `db:"serv_conf_id_fk" json:"servconfid"`
	SpId       int64 `db:"sp_id_fk" json:"spid"`
}

type DBServiceInstanceInsertRowModel struct {
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
	DBServiceInstanceDataModel
}

type DBServiceConfShortDataModel struct {
	ServConfId   int64  `db:"id" dbattr:"pri,auto"  json:"servconfid"`
	ServConfName string `db:"serv_conf_name" json:"servconfname"`
}

type DBServiceConfTemplateInsertDataModel struct {
	ServConfId int64 `db:"id" json:"servconfid"`
}

type DBServicePointConfigShortDataModel struct {
	SpId         int64   `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	SpName       string  `db:"sp_name" json:"spname"`
	SpcId        int64   `db:"spc_id_fk" json:"spcid"`
	SpcName      string  `db:"spc_name" json:"spcname"`
	ServConfId   *int64  `db:"serv_conf_id_fk" json:"servconfid"`
	ServConfName *string `db:"serv_conf_name" json:"servconfname"`
}
