package models

import (
	"time"

	gmodels "opensoach.com/models"
)

type PacketProccessExecution struct {
	Token          string
	DevicePacket   []byte
	InstanceDBConn string
	TokenInfo      *gmodels.DeviceTokenModel
}

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

type PacketVehicleTokenData struct {
	VehicleNo string `db:"vehicle_no" json:"vehicleno"`
}

type PacketVehicleDetailsData struct {
	VehicleNo string `db:"vehicle_no" json:"vehicleno"`
	Details   string `db:"details" json:"details"`
}

type PacketVhlTokenTxnData struct {
	TokenId           int64     `json:tokenid`
	ServiceInstanceID int64     `json:"servinid"`
	FOPCode           string    `json:"fopcode"`
	Status            int       `json:"status"`
	TxnData           string    `json:"txndata"`
	TxnDate           time.Time `json:"txndate"`
}

type PacketVhlTokenClaimData struct {
	TokenId int64  `json:tokenid`
	FOPCode string `json:"fopcode"`
}
