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
