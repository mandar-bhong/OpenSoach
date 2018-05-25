package models

import (
	"time"
)

type DBServiceInstanceDataTableRowModel struct {
	DBServiceInstanceDataRowModel
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBServiceInstanceDataRowModel struct {
	CpmId             int64     `db:"cpm_id_fk" json:"cpmid"`
	ServiceInstanceID int64     `db:"serv_in_id_fk" json:"servinid"`
	TransactionData   string    `db:"txn_data" json:"txndata"`
	TransactionDate   time.Time `db:"txn_date" json:"txndate"`
}
