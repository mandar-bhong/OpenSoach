package models

import (
	"time"
)

type DBMasterUserRowModel struct {
	ID           int64  `db:"id" json:"id"`
	UserName     string `db:"usr_name" json:"username"`
	Password     string `db:"usr_password" json:"password"`
	UserState    int    `db:"usr_state" json:"state"`
	UserCategory int    `db:"usr_category" json:"category"`
	UserRoleId   int64  `db:"urole_id_fk" json:"userroleid"`
}

type DBUserAuthInfo struct {
	ProdCode         string `db:"prod_code" json:"prod code"`
	Connectionstring string `db:"connection_string" json:"connectionstring"`
	CpmId            int64  `db:"cpm_id" json:"id"`
	CustomerId       int64  `db:"cust_id_fk" json:"customerid"`
	UserRoleCode     string `db:"urole_code" json:"userrolecode"`
}

type DBUserInfoMinDataModel struct {
	FirstName string `db:"fname" json:"fname"`
	LastName  string `db:"lname" json:"lname"`
}

type DBCustomerLoginInfoDataModel struct {
	CorpName string `db:"corp_name" json:"corpname"`
	CustName string `db:"cust_name" json:"custname"`
}

type DBCustomerInfoDataModel struct {
	CustId         int64     `db:"id" json:"custid"`
	CorpId         int64     `db:"corp_id_fk" json:"corpid"`
	CustName       string    `db:"cust_name" json:"custname"`
	CustState      int       `db:"cust_state" json:"custstate"`
	CustStateSince time.Time `db:"cust_state_since" json:"custstatesince"`
	CreatedOn      time.Time `db:"created_on" json:"createdon"`
	UpdateOn       time.Time `db:"updated_on" json:"updateon"`
}
