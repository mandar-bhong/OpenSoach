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

type DBDeviceFieldOperatorDataModel struct {
	DeviceId int64   `db:"dev_id_fk" json:"devid"`
	SpId     int64   `db:"sp_id_fk" json:"spid"`
	FopId    int64   `db:"fop_id_fk" json:"fopid"`
	Fopcode  string  `db:"fopcode" json:"fopcode"`
	FopName  *string `db:"fop_name" json:"fopname"`
	FopState int     `db:"fop_state" json:"fopstate"`
	FopArea  int     `db:"fop_area" json:"foparea"`
}

type DBFieldOperatorShortDataModel struct {
	FopId   int64   `db:"id" dbattr:"pri,auto"  json:"fopid"`
	Fopcode string  `db:"fopcode" json:"fopcode"`
	FopName *string `db:"fop_name" json:"fopname"`
}

type DBFopSpDataModel struct {
	FopId int64 `db:"fop_id_fk" dbattr:"pri"  json:"fopid"`
	SpId  int64 `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
}

type DBFopSpInsertRowModel struct {
	DBFopSpDataModel
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}

type DBFopSpAssociationDataModel struct {
	FopId  int64  `db:"fop_id_fk" dbattr:"pri"  json:"fopid"`
	SpId   int64  `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	SpName string `db:"sp_name" json:"spname"`
}
