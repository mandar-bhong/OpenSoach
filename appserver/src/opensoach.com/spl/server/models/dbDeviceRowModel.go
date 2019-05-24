package models

import (
	"time"
)

type DBDeviceConnectionStateUpdateRowModel struct {
	DeviceID        int64     `dbattr:"pri" db:"dev_id_fk" json:"deviceid"`
	ConnectionState int       `db:"connection_state" json:"connstate"`
	StateSince      time.Time `db:"connection_state_since" json:"connstatesince"`
}
