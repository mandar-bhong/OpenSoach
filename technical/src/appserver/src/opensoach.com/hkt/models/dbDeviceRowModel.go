package models

import "time"

type DBDeviceShortDataModel struct {
	DevId    int64  `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	DevName  string `db:"dev_name" json:"devname"`
	Serialno string `db:"serialno" json:"serialno"`
}

type DBDevStatusBatteryLevelUpdateDataModel struct {
	DevId             int64     `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	BatteryLevel      int       `db:"battery_level" json:"batterylevel"`
	BatteryLevelSince time.Time `db:"battery_level_since" json:"batterylevelsince"`
}

type DBDevStatusConnectionStateUpdateDataModel struct {
	DevId                int64     `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	ConnectionState      int       `db:"connection_state" json:"connectionstate"`
	ConnectionStateSince time.Time `db:"connection_state_since" json:"connectionstatesince"`
}
