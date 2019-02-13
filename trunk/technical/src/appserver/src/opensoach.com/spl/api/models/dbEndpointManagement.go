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

type DBDeviceUserListDataModel struct {
	UserId    int64  `json:"usr_id" db:"usr_id"`
	UserName  string `json:"usr_name" db:"usr_name"`
	UroleName string `json:"urole_name" db:"urole_name"`
	FirsName  string `json:"fname" db:"fname"`
	LastName  string `json:"lname" db:"lname"`
}
