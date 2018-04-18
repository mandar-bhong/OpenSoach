package models

import "time"

type DBSearchCustomerRequestFilterDataModel struct {
	CustName  *string `db:"cust_name" json:"custname"`
	CustState *int    `db:"cust_state" json:"custstate"`
}

type DBSearchUserRequestFilterDataModel struct {
	Id       *int64  `db:"id" json:"id"`
	Name     *string `db:"usr_name" json:"name"`
	Category *int    `db:"usr_category" json:"category"`
	State    *int    `db:"usr_state" json:"state"`
}

type DBSearchCorpRequestFilterDataModel struct {
	Name     *string `db:"corp_name" json:"name"`
	MobileNo *string `db:"corp_mobile_no" json:"mobileno"`
	EmailId  *string `db:"corp_email_id" json:"emailid"`
}

type DBSearchUserResponseFilterDataModel struct {
	Id            int64     `db:"id" json:"id"`
	UsrName       string    `db:"usr_name" json:"usrname"`
	UsrCategory   int       `db:"usr_category" json:"usrcategory"`
	UroleIdFk     *int64    `db:"urole_id_fk" json:"uroleidfk"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}

type DBSearchCorpResponseFilterDataModel struct {
	Id             int64     `dbattr:"pri,auto" db:"id" json:"id"`
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
	Poc1Name     string    `db:"poc1_name" json:"poc1name"`
	Poc1EmailID  string    `db:"poc1_email_id" json:"poc1emailid"`
	Poc1MobileNo string    `db:"poc1_mobile_no" json:"poc1mobileno"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdateOn     time.Time `db:"updated_on" json:"updateon"`
}
