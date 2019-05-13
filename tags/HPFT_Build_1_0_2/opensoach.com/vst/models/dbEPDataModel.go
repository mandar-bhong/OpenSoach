package models

import "time"

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

type DBEPSPVhlTokenDataModel struct {
	TokenId     int64     `db:"id" dbattr:"pri,auto"  json:"tokenid"`
	Token       int64     `db:"token" json:"token"`
	VhlId       int64     `db:"vhl_id_fk" json:"vhlid"`
	VehicleNo   string    `db:"vehicle_no" json:"vehicleno"`
	State       int64     `db:"state" json:"state"`
	GeneratedOn time.Time `db:"generated_on" json:"generatedon"`
}
