package models

type DBDeviceAuthInfoModel struct {
	CpmID            int64  `db:"id" json:"cpmid"`
	ServerAddress    string `db:"server_address" json:"serveraddress"`
	ConnectionString string `db:"connection_string" json:"connectionstring"`
}

type DBDeviceUserAuthInfoModel struct {
	Connectionstring string `db:"connection_string" json:"connectionstring"`
	ServerAddress    string `db:"server_address" json:"serveraddress"`
	CpmId            int64  `db:"cpm_id" json:"cpmid"`
	CustomerId       int64  `db:"cust_id_fk" json:"custid"`
	UserRoleId       int64  `db:"urole_id_fk" json:"uroleid"`
	UserRoleCode     string `db:"urole_code" json:"urolecode"`
}

type DBDeviceSharedUserAuthInfoModel struct {
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
