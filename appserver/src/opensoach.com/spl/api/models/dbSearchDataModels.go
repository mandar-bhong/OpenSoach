package models

import "time"

type DBSearchCustomerRequestFilterDataModel struct {
	CustName  *string `db:"cust_name" json:"custname"`
	CustState *int    `db:"cust_state" json:"custstate"`
	CorpID    *int64  `db:"corp_id_fk" json:"corpid"`
	ProdCode  *string `db:"prod_code" json:"prodcode"`
}

type DBSearchUserRequestFilterDataModel struct {
	CpmId       *int64  `db:"cpm_id_fk" json:"cpmid"`
	UsrName     *string `db:"usr_name" json:"usrname"`
	UsrCategory *int    `db:"usr_category" json:"usrcategory"`
	UsrState    *int    `db:"usr_state" json:"usrstate"`
}

type DBSearchCorpRequestFilterDataModel struct {
	CorpName     *string `db:"corp_name" json:"corpname"`
	CorpMobileNo *string `db:"corp_mobile_no" json:"corpmobileno"`
	CorpEmailId  *string `db:"corp_email_id" json:"corpemailid"`
}

type DBSearchUserResponseFilterDataModel struct {
	UserId        int64     `db:"id" json:"usrid"`
	UsrName       string    `db:"usr_name" json:"usrname"`
	Fname         *string   `db:"fname" json:"fname"`
	Lname         *string   `db:"lname" json:"lname"`
	Gender        *int      `db:"gender" json:"gender"`
	MobileNo      *string   `db:"mobile_no" json:"mobileno"`
	UsrCategory   int       `db:"usr_category" json:"usrcategory"`
	UroleId       *int64    `db:"urole_id_fk" json:"uroleid"`
	UroleName     *string   `db:"urole_name" json:"urolename"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}

type DBSearchCorpResponseFilterDataModel struct {
	CorpId         int64     `dbattr:"pri,auto" db:"id" json:"corpid"`
	CorpName       string    `db:"corp_name" json:"corpname"`
	CorpMobileNo   *string   `db:"corp_mobile_no" json:"corpmobileno"`
	CorpEmailId    *string   `db:"corp_email_id" json:"corpemailid"`
	CorpLandlineNo *string   `db:"corp_landline_no" json:"corplandlineno"`
	CreatedOn      time.Time `db:"created_on" json:"createdon"`
	UpdatedOn      time.Time `db:"updated_on" json:"updatedon"`
}

type DBSearchCustomerResponseFilterDataModel struct {
	CustID       int64     `db:"id" json:"custid"`
	CorpID       int64     `db:"corp_id_fk" json:"corpid"`
	CustName     string    `db:"cust_name" json:"custname"`
	CorpName     string    `db:"corp_name" json:"corpname"`
	CustState    int       `db:"cust_state" json:"custstate"`
	Poc1Name     *string   `db:"poc1_name" json:"poc1name"`
	Poc1EmailID  *string   `db:"poc1_email_id" json:"poc1emailid"`
	Poc1MobileNo *string   `db:"poc1_mobile_no" json:"poc1mobileno"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdateOn     time.Time `db:"updated_on" json:"updateon"`
}

type DBSearchDeviceRequestFilterDataModel struct {
	Serialno *string `db:"serialno" json:"serialno"`
	CustName *string `db:"cust_name" json:"custname"`
	CpmId    *int64  `dbattr:"pri,auto" db:"cpm.id" json:"cpmid"` //This is specical condition as query has multiple table join
	DevState *int    `db:"dev_state" json:"devstate"`
}

type DBSearchDeviceResponseFilterDataModel struct {
	DevId         int64     `dbattr:"pri,auto" db:"id" json:"devid"`
	CustId        *int64    `db:"cust_id_fk" json:"custid"`
	CustName      *string   `db:"cust_name" json:"custname"`
	Serialno      *string   `db:"serialno" json:"serialno"`
	DevState      int       `db:"dev_state" json:"devstate"`
	DevStateSince time.Time `db:"dev_state_since" json:"devstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}
