package models

type TaskAPICustProdAssociatedModel struct {
	CustId int64 `db:"cust_id_fk" json:"custid"`
	ProdId int64 `db:"prod_id_fk" json:"prodid"`
	DbiId  int64 `db:"dbi_id_fk" json:"dbiid"`
	CpmId  int64 `db:"id" json:"cpmid"`
}

type TaskDevProdAsscociatedModel struct {
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
	DevId int64 `db:"dev_id_fk" json:"devid"`
}
