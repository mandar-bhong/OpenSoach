package models

import "time"

type DBComplaintDataModel struct {
	SpId           int64   `db:"sp_id_fk" json:"spid"`
	ComplaintTitle string  `db:"complaint_title" json:"complainttitle"`
	Description    *string `db:"description" json:"description"`
	ComplaintBy    string  `db:"complaint_by" json:"complaintby"`
	MobileNo       *string `db:"mobile_no" json:"mobileno"`
	EmailId        *string `db:"email_id" json:"emailid"`
	EmployeeId     *string `db:"employee_id" json:"employeeid"`
	Severity       *int    `db:"severity" json:"severity"`
	ComplaintState int     `db:"complaint_state" json:"complaintstate"`
}

type DBComplaintInsertRowModel struct {
	DBComplaintDataModel
	CpmId    int64     `db:"cpm_id_fk" json:"cpmid"`
	RaisedOn time.Time `db:"raised_on" json:"raisedon"`
}

type DBComplaintUpdateRowModel struct {
	ComplaintId    int64      `db:"id" dbattr:"pri,auto"  json:"complaintid"`
	CpmId          int64      `db:"cpm_id_fk" json:"cpmid"`
	Severity       *int       `db:"severity" json:"severity"`
	ComplaintState int        `db:"complaint_state" json:"complaintstate"`
	Remarks        *string    `db:"remarks" json:"remarks"`
	ClosedOn       *time.Time `db:"closed_on" json:"closedon"`
}

type DBTopComplaintsFilterDataModel struct {
	SpId           *int64 `db:"sp_id_fk" json:"spid"`
	CpmId          int64  `db:"cpm_id_fk" json:"cpmid"`
	ComplaintState int    `db:"complaint_state" json:"complaintstate"`
}

type DBNoOfComplaintsPerMonthsFilterDataModel struct {
	SpId  *int64 `db:"sp_id_fk" json:"spid"`
	CpmId int64  `db:"cpm_id_fk" json:"cpmid"`
}

type DBNoOfComplaintsPerMonthDataModel struct {
	Year       int `db:"year" json:"year"`
	Month      int `db:"month" json:"month"`
	Open 	   int `db:"open" json:"open"`
	Closed     int `db:"closed" json:"closed"`
	InProgress int `db:"inprogress" json:"inprogress"`
}
