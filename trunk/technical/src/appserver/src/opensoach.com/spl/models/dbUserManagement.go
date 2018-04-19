package models

type DBUserAuthInfo struct {
	Connectionstring string `db:"connection_string" json:"connectionstring"`
	CpmId            int64  `db:"cpm_id" json:"id"`
	CustomerId       int64  `db:"cust_id_fk" json:"customerid"`
	UserRoleId       int64  `db:"urole_id_fk" json:"userroleid"`
	UserRoleCode     string `db:"urole_code" json:"userrolecode"`
}

type DBUserInfoMinDataModel struct {
	UserName  *string `db:"usr_name" json:"usrname"`
	FirstName *string `db:"fname" json:"fname"`
	LastName  *string `db:"lname" json:"lname"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

type DBCustomerLoginInfoDataModel struct {
	CorpName string `db:"corp_name" json:"corpname"`
	CustName string `db:"cust_name" json:"custname"`
}
