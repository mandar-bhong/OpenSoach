package models

import "time"

type DBSearchCustomerRequestFilterDataModel struct {
	Name  *string `db:"cust_name" json:"name"`
	State *int    `db:"cust_state" json:"state"`
}

type DBSearchUserRequestFilterDataModel struct {
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
	Id            int64     `dbattr:"pri,auto" db:"id" json:"id"`
	UsrName       string    `db:"usr_name" json:"usrname"`
	UsrCategory   int       `db:"usr_category" json:"usrcategory"`
	UroleIdFk     *int64    `db:"urole_id_fk" json:"uroleidfk"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}
