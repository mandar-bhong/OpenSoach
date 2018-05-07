package models

type DBFieldOperatorDataModel struct {
	Fopcode   string  `db:"fopcode" json:"fopcode"`
	FopName   *string `db:"fop_name" json:"fopname"`
	MobileNo  string  `db:"mobile_no" json:"mobileno"`
	EmailId   *string `db:"email_id" json:"emailid"`
	ShortDesc *string `db:"short_desc" json:"shortdesc"`
	FopState  int     `db:"fop_state" json:"fopstate"`
	FopArea   int     `db:"fop_area" json:"foparea"`
}

type DBFieldOperatorRowModel struct {
	CpmId int64 `db:"cpm_id_fk" json:"cmpid"`
	DBFieldOperatorDataModel
}

type DBFieldOperatorUpdateRowModel struct {
	FopId     int64   `db:"id" dbattr:"pri,auto"  json:"fopid"`
	CpmId     int64   `db:"cpm_id_fk" json:"cpmid"`
	FopName   *string `db:"fop_name" json:"fopname"`
	MobileNo  string  `db:"mobile_no" json:"mobileno"`
	EmailId   *string `db:"email_id" json:"emailid"`
	ShortDesc *string `db:"short_desc" json:"shortdesc"`
	FopState  int     `db:"fop_state" json:"fopstate"`
	FopArea   int     `db:"fop_area" json:"foparea"`
}
