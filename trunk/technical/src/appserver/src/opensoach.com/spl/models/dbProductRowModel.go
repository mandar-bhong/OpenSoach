package models

type DBProductBriefRowModel struct {
	CPMID        int64  `db:"cpm_id" json:"cpmid"`
	ProductCode  string `db:"prod_code" json:"prodcode"`
	CustomerID   int64  `db:"cust_id_fk" json:"custid"`
	CustomerName string `db:"cust_name" json:"custname"`
	DBConnection string `db:"cust_name" json:"dbconn"`
}

type DBProductRowModel struct {
	CPMID        int64  `db:"cpm_id" json:"cpmid"`
	ProductCode  string `db:"prod_code" json:"prodcode"`
	CustomerID   int64  `db:"cust_id_fk" json:"custid"`
	CustomerName string `db:"cust_name" json:"custname"`
}
