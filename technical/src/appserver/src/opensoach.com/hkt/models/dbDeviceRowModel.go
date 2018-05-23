package models

type DBDeviceShortDataModel struct {
	DevId   int64  `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	DevName string `db:"dev_name" json:"devname"`
}
