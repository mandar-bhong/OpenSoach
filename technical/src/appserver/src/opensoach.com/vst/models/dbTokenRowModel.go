package models

import "time"

type DBTokenDataModel struct {
	Token          int64  `db:"token" json:"token"`
	VhlId          int64  `db:"vhl_id_fk" json:"vhlid"`
	MappingDetails string `db:"mapping_details" json:"mappingdetails"`
	State          int64  `db:"state" json:"state"`
}

type DBTokenInsertRowModel struct {
	DBTokenDataModel
	GeneratedOn time.Time `db:"generated_on" json:"generatedon"`
}

type DBTokenMappingDetailsUpdateModel struct {
	TokenId        int64  `db:"id" dbattr:"pri,auto"  json:"tokenid"`
	MappingDetails string `db:"mapping_details" json:"mappingdetails"`
	State          int64  `db:"state" json:"state"`
}

type TokenMappingDetailsModel struct {
	TokenConfigId int64 `json:"tokenconfigid"`
	JobCreationId int64 `json:"jobcreationid"`
	JobExeId      int64 `json:"jobexeid"`
}

type DBDeviceVhlTokenModel struct {
	DeviceId    int64     `db:"dev_id_fk" json:"devid"`
	SpId        int64     `db:"sp_id_fk" json:"spid"`
	TokenId     int64     `db:"id" dbattr:"pri,auto"  json:"tokenid"`
	Token       int64     `db:"token" json:"token"`
	VhlId       int64     `db:"vhl_id_fk" json:"vhlid"`
	VehicleNo   string    `db:"vehicle_no" json:"vehicleno"`
	State       int64     `db:"state" json:"state"`
	GeneratedOn time.Time `db:"generated_on" json:"generatedon"`
}

type DBTokenStateUpdateModel struct {
	TokenId int64 `db:"id" dbattr:"pri,auto"  json:"tokenid"`
	State   int64 `db:"state" json:"state"`
}
