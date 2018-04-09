package models

type DBMasterUserRowModel struct {
	ID           int64  `db:"id" json:"id"`
	UserName     string `db:"usr_name" json:"username"`
	Password     string `db:"usr_password" json:"password"`
	UserState    int    `db:"usr_state" json:"state"`
	UserCategory int    `db:"usr_category" json:"category"`
	UserRoleId   int    `db:"urole_id_fk" json:"userroleid"`
}

type DBUserAuthInfo struct {
	ProdCode         string `db:"prod_code" json:"prod code"`
	Connectionstring string `db:"connection_string" json:"connectionstring"`
	CpmId            int    `db:"cpm_id" json:"id"`
	CustomerId       int    `db:"cust_id_fk" json:"customerid"`
	UserRoleCode     string `db:"urole_code" json:"userrolecode"`
}
