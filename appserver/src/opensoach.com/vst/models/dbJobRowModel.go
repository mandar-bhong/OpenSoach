package models

import "time"

type DBJobStatusUpdateRowModel struct {
	TokenId int64 `db:"id" json:"tokenid"`
	State   int64 `db:"state" json:"state"`
}

type DBJobDeliveredTxnRowModel struct {
	CpmId             int64     `db:"cpm_id_fk" json:"cpmid"`
	ServiceInstanceID int64     `db:"serv_in_id_fk" json:"servinid"`
	FOPCode           string    `db:"fopcode" json:"fopcode"`
	Status            int       `db:"status" json:"status"`
	TransactionData   string    `db:"txn_data" json:"txndata"`
	TransactionDate   time.Time `db:"txn_date" json:"txndate"`
}

type DBJobDeliveredTxnDataModel struct {
	Tokenid      int64 `db:"tokenid" json:"tokenid"`
	BilledAmount int64 `db:"billedamount" json:"billedamount"`
}

type DBJobVhlInfoModel struct {
	VehicleNo string `db:"vehicle_no" json:"vehicleno"`
	MobileNo  string `db:"mobileno" json:"mobileno"`
	Name      string `db:"name" json:"name"`
	Kms       string `db:"kms" json:"kms"`
	Petrol    string `db:"petrol" json:"petrol"`
}
