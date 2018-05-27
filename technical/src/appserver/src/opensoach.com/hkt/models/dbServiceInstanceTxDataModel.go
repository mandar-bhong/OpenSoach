package models

import (
	"time"
)

type DBServiceInstanceTxDataTableRowModel struct {
	DBServiceInstanceTxDataRowModel
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBServiceInstanceTxDataRowModel struct {
	CpmId             int64     `db:"cpm_id_fk" json:"cpmid"`
	ServiceInstanceID int64     `db:"serv_in_id_fk" json:"servinid"`
	FOPCode           string    `db:"fopcode" json:"fopcode"`
	Status            int       `db:"status" json:"status"`
	TransactionData   string    `db:"txn_data" json:"txndata"`
	TransactionDate   time.Time `db:"txn_date" json:"txndate"`
}
