package models

type DBDeviceAuthInfoModel struct {
	CpmID         int64  `db:"id" json:"cpmid"`
	ServerAddress string `db:"server_address" json:"serveraddress"`
}
