package models

type DBSearchFieldOperatorRequestFilterDataModel struct {
	Fopcode  *string `db:"fopcode" json:"fopcode"`
	FopName  *string `db:"fop_name" json:"fopname"`
	MobileNo *string `db:"mobile_no" json:"mobileno"`
	EmailId  *string `db:"email_id" json:"emailid"`
}

type DBSearchFieldOperatorResponseFilterDataModel struct {
	FopId    int64   `db:"id" dbattr:"pri,auto"  json:"fopid"`
	Fopcode  string  `db:"fopcode" json:"fopcode"`
	FopName  *string `db:"fop_name" json:"fopname"`
	MobileNo string  `db:"mobile_no" json:"mobileno"`
	EmailId  *string `db:"email_id" json:"emailid"`
	FopState int     `db:"fop_state" json:"fopstate"`
	FopArea  int     `db:"fop_area" json:"foparea"`
}
