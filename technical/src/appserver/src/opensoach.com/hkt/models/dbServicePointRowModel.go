package models

import "time"

type DBSpUpdateRowModel struct {
	SpId         int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	SpName       string    `db:"sp_name" json:"spname"`
	ShortDesc    *string   `db:"short_desc" json:"shortdesc"`
	SpState      int       `db:"sp_state" json:"spstate"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
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

type DBSpDataRowModel struct {
	SpId      int64   `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	SpcId     int64   `db:"spc_id_fk" json:"spcid"`
	SpName    string  `db:"sp_name" json:"spname"`
	ShortDesc *string `db:"short_desc" json:"shortdesc"`
	SpState   int     `db:"sp_state" json:"spstate"`
}

type DBSpInsertRowModel struct {
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
	DBSpDataRowModel
}

type DBSpCategoryShortDataModel struct {
	SpcId   int64  `db:"id" dbattr:"pri,auto"  json:"spcid"`
	SpcName string `db:"spc_name" json:"spcname"`
}

type DBDevSpMappingDataModelModel struct {
	DevId int64 `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	SpId  int64 `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
}

type DBDevSpMappingInsertRowModel struct {
	DBDevSpMappingDataModelModel
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}

type DBFopSpAssociationDataModel struct {
	FopId  int64  `db:"fop_id_fk" dbattr:"pri"  json:"fopid"`
	SpId   int64  `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	SpName string `db:"sp_name" json:"spname"`
}

type DBServicePointShortDataModel struct {
	SpId   int64  `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	SpName string `db:"sp_name" json:"spname"`
}
