package models

import "time"

type DBCustProdMappingInsertRowModel struct {
	CustId        int64     `db:"cust_id_fk" json:"custid"`
	ProdId        int64     `db:"prod_id_fk" json:"prodid"`
	DbiId         int64     `db:"dbi_id_fk" json:"dbiid"`
	CpmState      int       `db:"cpm_state" json:"cpmstate"`
	CpmStateSince time.Time `db:"cpm_state_since" json:"cpmstatesince"`
}

type DBCustProdAssociationInfoRowModel struct {
	CpmId    int64  `dbattr:"pri,auto" db:"id" json:"cpmid"`
	CpmState int    `db:"cpm_state" json:"cpmstate"`
	ProdId   int64  `db:"prod_id_fk" json:"prodid"`
	DbiId    int64  `db:"dbi_id_fk" json:"dbiid"`
	ProdCode string `db:"prod_code" json:"prodcode"`
	DbiName  string `db:"dbi_name" json:"dbiname"`
}

type DBCpmStateUpdateRowModel struct {
	CpmId         int64     `dbattr:"pri,auto" db:"id" json:"cpmid"`
	CpmState      int       `db:"cpm_state" json:"cpmstate"`
	CpmStateSince time.Time `db:"cpm_state_since" json:"cpmstatesince"`
}

type DBCustomerUpdateRowModel struct {
	CustId         int64     `dbattr:"pri,auto" db:"id" json:"custid"`
	CustName       string    `db:"cust_name" json:"custname"`
	CustState      int       `db:"cust_state" json:"custstate"`
	CustStateSince time.Time `db:"cust_state_since" json:"custstatesince"`
}

type DBCustShortDataModel struct {
	CustId   int64  `dbattr:"pri,auto" db:"id" json:"custid"`
	CustName string `db:"cust_name" json:"custname"`
}

type DBCustSpDataModel struct {
	CpmId    int64  `db:"id" json:"cpmid"`
	ProdCode string `db:"prod_code" json:"prodcode"`
	SpCount  int    `db:"count" json:"spcount"`
}

type DBServicepointInsertRowModel struct {
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpState      int       `db:"sp_state" json:"spstate"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
}
