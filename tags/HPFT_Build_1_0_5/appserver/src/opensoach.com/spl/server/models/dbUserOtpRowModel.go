package models

type DBUserOtpRowModel struct {
	UsrName string `db:"usr_name" json:"usrname"`
	Otp     string `db:"otp" json:"otp"`
}
