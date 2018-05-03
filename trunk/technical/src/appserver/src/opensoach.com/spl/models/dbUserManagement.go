package models

import "time"

type DBUserAuthInfo struct {
	Connectionstring string `db:"connection_string" json:"connectionstring"`
	CpmId            int64  `db:"cpm_id" json:"cpmid"`
	CustomerId       int64  `db:"cust_id_fk" json:"custid"`
	UserRoleId       int64  `db:"urole_id_fk" json:"uroleid"`
	UserRoleCode     string `db:"urole_code" json:"urolecode"`
}

type DBUserInfoMinDataModel struct {
	UserName  *string `db:"usr_name" json:"usrname"`
	FirstName *string `db:"fname" json:"fname"`
	LastName  *string `db:"lname" json:"lname"`
}

type DBCustomerLoginInfoDataModel struct {
	CorpName string `db:"corp_name" json:"corpname"`
	CustName string `db:"cust_name" json:"custname"`
}

type DBSplMasterUserRowModel struct {
	UsrId         int64     `dbattr:"pri,auto" db:"id" json:"usrid"`
	UsrName       string    `db:"usr_name" json:"usrname"`
	UsrPassword   string    `db:"usr_password" json:"usrpassword"`
	UsrCategory   int       `db:"usr_category" json:"usrcategory"`
	UroleId       *int64    `db:"urole_id_fk" json:"uroleid"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
}

type DBSplMasterUsrDetailsRowModel struct {
	UsrId              int64   `dbattr:"pri" db:"usr_id_fk" json:"usrid"`
	Fname              *string `db:"fname" json:"fname"`
	Lname              *string `db:"lname" json:"lname"`
	Gender             *int    `db:"gender" json:"gender"`
	MobileNo           *string `db:"mobile_no" json:"mobileno"`
	AlternateContactNo *string `db:"alternate_contact_no" json:"alternatecontactno"`
}

type DBUsrCpmRowModel struct {
	UserId  int64 `db:"user_id_fk" json:"userid"`
	CpmId   int64 `db:"cpm_id_fk" json:"cpmid"`
	UroleId int64 `db:"urole_id_fk" json:"uroleid"`
}

type DBUsrCpmUpdateRowModel struct {
	UcpmId int64 `dbattr:"pri,auto" db:"id" json:"ucpmid"`
	DBUsrCpmRowModel
}

type DBUroleShortDataModel struct {
	UroleId   int64  `dbattr:"pri,auto" db:"id" json:"uroleid"`
	UroleCode string `db:"urole_code" json:"urolecode"`
	UroleName string `db:"urole_name" json:"urolename"`
}
