package models

type DBUserAuthInfo struct {
	Connectionstring string `db:"connection_string" json:"connectionstring"`
	CpmId            int64  `db:"cpm_id" json:"id"`
	CustomerId       int64  `db:"cust_id_fk" json:"customerid"`
	UserRoleCode     string `db:"urole_code" json:"userrolecode"`
}

type DBUserInfoMinDataModel struct {
	FirstName *string `db:"fname" json:"fname"`
	LastName  *string `db:"lname" json:"lname"`
}

type DBCustomerLoginInfoDataModel struct {
	CorpName string `db:"corp_name" json:"corpname"`
	CustName string `db:"cust_name" json:"custname"`
}
