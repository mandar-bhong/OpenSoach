package models

import (
	"time"
)

type PacketServiceInstanceData struct {
	ServiceInstanceID int64     `json:"servinid"`
	FOPCode           string    `json:"fopcode"`
	Status            int       `json:"status"`
	TxnData           string    `json:"txndata"`
	TxnDate           time.Time `json:"txndate"`
}

type PacketComplaintData struct {
	ComplaintTitle string    `db:"complaint_title" json:"complainttitle"`
	Description    *string   `db:"description" json:"description"`
	ComplaintBy    string    `db:"complaint_by" json:"complaintby"`
	MobileNo       *string   `db:"mobile_no" json:"mobileno"`
	EmailId        *string   `db:"email_id" json:"emailid"`
	EmployeeId     *string   `db:"employee_id" json:"employeeid"`
	RaisedOn       time.Time `db:"raised_on" json:"raisedon"`
}

type PacketFeedbackData struct {
	Feedback        int       `jason:"feedback"`
	FeedbackComment *string   `json:"comment"`
	RaisedOn        time.Time `db:"raised_on" json:"raisedon"`
}

type PacketDeviceBatteryLevelUpdateData struct {
	BatteryLevel int `db:"battery_level" json:"batterylevel"`
}
