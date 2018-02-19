package models

type DBMasterUserRowModel struct {
	ID           int64  `db:"id" json:"id"`
	UserName     string `db:"usr_name" json:"username"`
	Password     string `db:"usr_password" json:"password"`
	UserState    int    `db:"usr_state" json:"state"`
	UserCategory int    `db:"usr_category" json:"category"`
}
