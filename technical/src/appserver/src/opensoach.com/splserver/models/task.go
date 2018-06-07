package models

import "time"

type APITaskHandlerFunc func(msg string, sessionkey string, token string, taskData interface{}) (error, APITaskResultModel)

type APITaskHandlerModel struct {
	Handler     APITaskHandlerFunc
	PayloadType interface{}
}

type APITaskResultModel struct {
	IsSuccess bool                        `json:"isSuceess"`
	Data      interface{}                 `json:"data"`
	ErrorData APITaskResultErrorDataModel `json:"errordata"`
}

type APITaskResultErrorDataModel struct {
	ErrorCode int         `json:"code"`
	Data      interface{} `json:"data"`
}

type APITaskDBInstanceCpmIdInsertModel struct {
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}

type APITaskDBInstanceDevInsertRowModel struct {
	DevId    int64  `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	CpmId    int64  `db:"cpm_id_fk" json:"cpmid"`
	DevName  string `db:"dev_name" json:"devname"`
	Serialno string `db:"serialno" json:"serialno"`
}

type APITaskDBInstanceSpCategoryInsertModel struct {
	SpcId     int64   `db:"id" dbattr:"pri,auto"  json:"spcid"`
	CpmId     int64   `db:"cpm_id_fk" json:"cpmid"`
	SpcName   string  `db:"spc_name" json:"spcname"`
	ShortDesc *string `db:"short_desc" json:"shortdesc"`
}

type APITaskDBNodeSpInsertRowModel struct {
	SpId         int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	SpName       string    `db:"sp_name" json:"spname"`
	SpState      int       `db:"sp_state" json:"spstate"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
}

type APITaskDBInstanceDevStatusInsertModel struct {
	DevId                int64     `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	ConnectionState      int       `db:"connection_state" json:"connectionstate"`
	ConnectionStateSince time.Time `db:"connection_state_since" json:"connectionstatesince"`
	SyncState            int       `db:"sync_state" json:"syncstate"`
	SyncStateSince       time.Time `db:"sync_state_since" json:"syncstatesince"`
	BatteryLevel         int       `db:"battery_level" json:"batterylevel"`
	BatteryLevelSince    time.Time `db:"battery_level_since" json:"batterylevelsince"`
}
