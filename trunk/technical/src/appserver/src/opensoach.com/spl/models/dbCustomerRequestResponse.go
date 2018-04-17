package models

import (
	"time"
)

type CustomerAddRequest struct {
	CorporationID      int64     `db:"corp_id"   json:"corpid"`
	CustomerName       string    `db:"cust_name" json:"custname"`
	CustomerState      int       `db:"cust_state" json:"custstate"`
	CustomerStateSince time.Time `db:"cust_state_since" json:"custstatesince"`
}
