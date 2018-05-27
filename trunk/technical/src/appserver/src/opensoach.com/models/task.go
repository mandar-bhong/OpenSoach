package models

import "time"

type APITaskProcessInterface interface {
	GetCPM() int64
}

type APITaskProcessModel struct {
	TaskToken    string      `json:"token"`
	SessionToken string      `json:"sessionkey"`
	TaskTag      string      `json:"tag"`
	TaskPayload  interface{} `json:"payload"`
}

type TaskExeContextModel struct {
	CustId int64  `db:"cust_id_fk" json:"custid"`
	ProdId int64  `db:"prod_id_fk" json:"prodid"`
	DBConn string `db:"prod_id_fk" json:"dbconn"`
}

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

type TaskCustServicePointAssociatedModel struct {
	SpIdList     []int64   `json:"spidlist"`
	CpmId        int64     `json:"cpmid"`
	SpState      int       `json:"spstate"`
	SpStateSince time.Time `json:"spstatesince"`
}
