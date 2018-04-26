package models

import "time"

type DBCustProdMappingInsertRowModel struct {
	CustId        int64     `db:"cust_id_fk" json:"custid"`
	ProdId        int64     `db:"prod_id_fk" json:"prodid"`
	DbiId         int64     `db:"dbi_id_fk" json:"dbiid"`
	CpmState      int       `db:"cpm_state" json:"cpmstate"`
	CpmStateSince time.Time `db:"cpm_state_since" json:"cpmstatesince"`
}
