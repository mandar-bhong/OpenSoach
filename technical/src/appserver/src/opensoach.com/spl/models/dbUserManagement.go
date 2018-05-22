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
	UserId         int64     `db:"user_id_fk" json:"userid"`
	CpmId          int64     `db:"cpm_id_fk" json:"cpmid"`
	UroleId        int64     `db:"urole_id_fk" json:"uroleid"`
	UcpmState      int       `db:"ucpm_state" json:"ucpmstate"`
	UcpmStateSince time.Time `db:"ucpm_state_since" json:"ucpmstatesince"`
}

type DBUsrCpmStateUpdateRowModel struct {
	UcpmId         int64     `dbattr:"pri,auto" db:"id" json:"ucpmid"`
	UcpmState      int       `db:"ucpm_state" json:"ucpmstate"`
	UcpmStateSince time.Time `db:"ucpm_state_since" json:"ucpmstatesince"`
}

type DBUroleShortDataModel struct {
	UroleId   int64   `dbattr:"pri,auto" db:"id" json:"uroleid"`
	UroleCode string  `db:"urole_code" json:"urolecode"`
	UroleName string  `db:"urole_name" json:"urolename"`
	ProdCode  *string `db:"prod_code" json:"prodcode"`
}

type DBUserProdAssociationDataModel struct {
	UcpmId         int64      `dbattr:"pri,auto" db:"id" json:"ucpmid"`
	CustName       *string    `db:"cust_name" json:"custname"`
	ProdCode       *string    `db:"prod_code" json:"prodcode"`
	UroleCode      *string    `db:"urole_code" json:"urolecode"`
	UcpmState      *int       `db:"ucpm_state" json:"ucpmstate"`
	UcpmStateSince *time.Time `db:"ucpm_state_since" json:"ucpmstatesince"`
}

type DBUserUpdateRowModel struct {
	UserId        int64     `dbattr:"pri,auto" db:"id" json:"userid"`
	UroleId       *int64    `db:"urole_id_fk" json:"uroleid"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
}

type DBCUUserUpateRowModel struct {
	UserId        int64     `dbattr:"pri,auto" db:"id" json:"userid"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
}

type DBCUUcpmUpdateRowModel struct {
	UserId  int64 `db:"user_id_fk" json:"userid"`
	UroleId int64 `db:"urole_id_fk" json:"uroleid"`
}

type DBUserInfoDataModel struct {
	UserId        int64     `dbattr:"pri,auto" db:"id" json:"userid"`
	UsrName       string    `db:"usr_name" json:"usrname"`
	UsrCategory   int       `db:"usr_category" json:"usrcategory"`
	UroleId       *int64    `db:"urole_id_fk" json:"uroleid"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}

type DBCUUserInfoDataModel struct {
	UserId        int64     `dbattr:"pri,auto" db:"id" json:"userid"`
	UsrName       string    `db:"usr_name" json:"usrname"`
	UsrCategory   int       `db:"usr_category" json:"usrcategory"`
	UroleId       int64     `db:"urole_id_fk" json:"uroleid"`
	UroleName     string    `db:"urole_name" json:"urolename"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}
