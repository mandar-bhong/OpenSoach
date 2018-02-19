package models

type DBMasterUserRowModel struct {
	ID           string `db:"id" json:"id"`
	UserName     string `db:"usr_name" json:"username"`
	Password     string `db:"usr_password" json:"password"`
	UserState    string `db:"usr_state" json:"state"`
	UserCategory string `db:"usr_category" json:"category"`
}
