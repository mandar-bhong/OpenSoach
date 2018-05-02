package models

type APIFieldOperatorAddRequest struct {
	Fopcode   string  `db:"fopcode" json:"fopcode"`
	FopName   *string `db:"fop_name" json:"fopname"`
	MobileNo  string  `db:"mobile_no" json:"mobileno"`
	EmailId   *string `db:"email_id" json:"emailid"`
	ShortDesc *string `db:"short_desc" json:"shortdesc"`
	FopState  int     `db:"fop_state" json:"fopstate"`
	FopArea   int     `db:"fop_area" json:"foparea"`
}
