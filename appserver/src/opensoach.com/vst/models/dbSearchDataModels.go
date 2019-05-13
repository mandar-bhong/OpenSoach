package models

import "time"

type DBSearchFieldOperatorRequestFilterDataModel struct {
	Fopcode  *string `db:"fopcode" json:"fopcode"`
	FopName  *string `db:"fop_name" json:"fopname"`
	MobileNo *string `db:"mobile_no" json:"mobileno"`
	EmailId  *string `db:"email_id" json:"emailid"`
	CpmId    int64   `db:"cpm_id_fk" json:"cpmid"`
	FopState *int    `db:"fop_state" json:"fopstate"`
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
	SpId           *int64  `db:"sp.sp_id_fk" json:"spid"`
	CpmId          int64   `db:"sp.cpm_id_fk" json:"cpmid"`
}

type DBSearchComplaintResponseFilterDataModel struct {
	ComplaintId    int64      `db:"id" dbattr:"pri,auto"  json:"complaintid"`
	SpId           int64      `db:"sp_id_fk"  json:"spid"`
	SpName         string     `db:"sp_name" json:"spname"`
	ComplaintTitle string     `db:"complaint_title" json:"complainttitle"`
	Description    *string    `db:"description" json:"description"`
	ComplaintBy    string     `db:"complaint_by" json:"complaintby"`
	Severity       *int       `db:"severity" json:"severity"`
	RaisedOn       time.Time  `db:"raised_on" json:"raisedon"`
	ComplaintState int        `db:"complaint_state" json:"complaintstate"`
	ClosedOn       *time.Time `db:"closed_on" json:"closedon"`
}

type DBSearchServiceConfRequestFilterModel struct {
	ConfTypeCode *string `db:"conf_type_code" json:"conftypecode"`
	ServConfName *string `db:"serv_conf_name" json:"servconfname"`
	CpmId        int64   `db:"cpm_id_fk" json:"cpmid"`
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
	CpmId   int64   `db:"serv_conf_in.cpm_id_fk" json:"cpmid"`
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
	CpmId   *int64  `db:"sp.cpm_id_fk" json:"cpmid"`
}

type DBSearchServicePointResponseFilterDataModel struct {
	SpId         int64     `db:"sp_id_fk" json:"spid"`
	SpName       string    `db:"sp_name" json:"spname"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	SpcName      string    `db:"spc_name" json:"spcname"`
	DevId        *int64    `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	DevName      *string   `db:"dev_name" json:"devname"`
	ServConfId   *int64    `db:"serv_conf_id_fk" json:"servconfid"`
	SpState      int       `db:"sp_state" json:"spstate"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
}

type DBSearchDeviceRequestFilterDataModel struct {
	Serialno        *string `db:"serialno" json:"serialno"`
	DevName         *string `db:"dev_name" json:"devname"`
	ConnectionState *int    `db:"connection_state" json:"connectionstate"`
	CpmId           int64   `db:"cpm_id_fk" json:"cpmid"`
}

type DBSearchDeviceResponseFilterDataModel struct {
	DevId                int64      `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	DevName              *string    `db:"dev_name" json:"devname"`
	Serialno             string     `db:"serialno" json:"serialno"`
	CreatedOn            time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn            time.Time  `db:"updated_on" json:"updatedon"`
	ConnectionState      *int       `db:"connection_state" json:"connectionstate"`
	ConnectionStateSince *time.Time `db:"connection_state_since" json:"connectionstatesince"`
	SyncState            *int       `db:"sync_state" json:"syncstate"`
	SyncStateSince       *time.Time `db:"sync_state_since" json:"syncstatesince"`
	BatteryLevel         *int       `db:"battery_level" json:"batterylevel"`
	BatteryLevelSince    *time.Time `db:"battery_level_since" json:"batterylevelsince"`
}

type DBSearchFeedbackRequestFilterDataModel struct {
	SpId     *int64 `db:"sp_id_fk" json:"spid"`
	Feedback *int   `db:"feedback" json:"feedback"`
	CpmId    int64  `db:"cpm_id_fk" json:"cpmid"`
}

type DBSearchFeedbackResponseFilterDataModel struct {
	FeedbackId      int64   `db:"id" dbattr:"pri,auto"  json:"feedbackid"`
	Feedback        int     `db:"feedback" json:"feedback"`
	FeedbackComment *string `db:"feedback_comment" json:"feedbackcomment"`
}

type DBSearchJobRequestFilterDataModel struct {
	VehicleNo *string    `db:"vehicle_no" json:"vehicleno"`
	Status    *int64     `db:"status" json:"status"`
	StartDate *time.Time `json:"startdate"`
	EndDate   *time.Time `json:"enddate"`
}

type DBSearchJobResponseFilterDataModel struct {
	VehicleId   int64      `db:"vehicleid" json:"vehicleid"`
	VehicleNo   string     `db:"vehicle_no" json:"vehicleno"`
	GeneratedOn time.Time  `db:"generated_on" json:"generatedon"`
	Token       int64      `db:"token" json:"token"`
	TokenId     int64      `db:"tokenid" json:"tokenid"`
	State       int64      `db:"state" json:"state"`
	Intime      time.Time  `db:"intime" json:"intime"`
	Outtime     *time.Time `db:"outtime" json:"outtime"`
}
