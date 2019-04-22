package models

type DBUserInfoModel struct {
	UsrId     int64   `dbattr:"pri,auto" db:"id" json:"usrid"`
	UsrName   string  `db:"usr_name" json:"usrname"`
	UroleCode string  `db:"urole_code" json:"urolecode"`
	UroleName string  `db:"urole_name" json:"urolename"`
	Fname     *string `db:"fname" json:"fname"`
	Lname     *string `db:"lname" json:"lname"`
}
