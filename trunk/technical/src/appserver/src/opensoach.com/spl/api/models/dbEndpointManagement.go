package models

type DBDeviceAuthInfoModel struct {
	CpmID         int64  `db:"id" json:"cpmid"`
	ServerAddress string `db:"server_address" json:"serveraddress"`
}

type DBDeviceUserAuthInfoModel struct {
	UserId    int64  `db:"id" json:"userid"`
	UserName  string `db:"usr_name" json:"username"`
	FirstName string `db:"fname" json:"firstname"`
	LastName  string `db:"lname" json:"lastname"`
}
