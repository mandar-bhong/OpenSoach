package models

type DBUserActivationRowModel struct {
	Id              int    `dbattr:"pri" db:"id" json:"activationid"`
	UserId          int64  `db:"usr_id_fk" json:"userid"`
	Code            string `db:"code" json:"code"`
	PasswordChanged bool   `db:"password_changed" json:"passwordchanged"`
}
