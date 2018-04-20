package models

type DBSplCorpRowModel struct {
	CorpId         int64   `dbattr:"pri,auto" db:"id" json:"corpid"`
	CorpName       string  `db:"corp_name" json:"corpname"`
	CorpMobileNo   *string `db:"corp_mobile_no" json:"corpmobileno"`
	CorpEmailId    *string `db:"corp_email_id" json:"corpemailid"`
	CorpLandlineNo *string `db:"corp_landline_no" json:"corplandlineno"`
}

type DBCorpShortDataModel struct {
	CorpID   int64  `db:"id" json:"corpid"`
	CorpName string `db:"corp_name" json:"corpname"`
}
