package models

import "time"

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

type DBSearchComplaintRequestFilterDataModel struct {
	ComplaintTitle *string `db:"complaint_title" json:"complainttitle"`
	ComplaintState *int    `db:"complaint_state" json:"complaintstate"`
}

type DBSearchComplaintResponseFilterDataModel struct {
	ComplaintId    int64      `db:"id" dbattr:"pri,auto"  json:"complaintid"`
	SpName         string     `db:"sp_name" json:"spname"`
	ComplaintTitle string     `db:"complaint_title" json:"complainttitle"`
	Description    *string    `db:"description" json:"description"`
	ComplaintBy    string     `db:"complaint_by" json:"complaintby"`
	Severity       int        `db:"severity" json:"severity"`
	RaisedOn       time.Time  `db:"raised_on" json:"raisedon"`
	ComplaintState int        `db:"complaint_state" json:"complaintstate"`
	ClosedOn       *time.Time `db:"closed_on" json:"closedon"`
}

type DBSearchServiceConfRequestFilterModel struct {
	ConfTypeCode *string `db:"conf_type_code" json:"conftypecode"`
	ServConfName *string `db:"serv_conf_name" json:"servconfname"`
}

type DBSearchServiceConfResponseFilterModel struct {
	ServConfId   int64     `db:"id" dbattr:"pri,auto"  json:"servconfid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	ConfTypeCode string    `db:"conf_type_code" json:"conftypecode"`
	ServConfName string    `db:"serv_conf_name" json:"servconfname"`
	ShortDesc    *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSearchServiceInstanceRequestFilterModel struct {
	SpcName *string `db:"spc_name" json:"spcname"`
}

type DBSearchServiceInstanceResponseFilterModel struct {
	SpId         int64  `db:"sp_id_fk" json:"spid"`
	ServInId     int64  `db:"id" dbattr:"pri,auto"  json:"servinid"`
	ConfTypeCode string `db:"conf_type_code" json:"conftypecode"`
	ServConfName string `db:"serv_conf_name" json:"servconfname"`
}

type DBSearchServicePointRequestFilterDataModel struct {
	SpName  *string `db:"sp_name" json:"spname"`
	SpcId   *int64  `db:"spc_id_fk" json:"spcid"`
	SpState *int    `db:"sp_state" json:"spstate"`
	DevId   *int64  `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
}

type DBSearchServicePointResponseFilterDataModel struct {
	SpId         int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	SpName       string    `db:"sp_name" json:"spname"`
	ShortDesc    *string   `db:"short_desc" json:"shortdesc"`
	SpState      int       `db:"sp_state" json:"spstate"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
}
