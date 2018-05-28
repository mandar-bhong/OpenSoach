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
	SpId        int64   `db:"sp_id_fk" json:"spid"`
	Description *string `db:"description" json:"description"`
	ComplaintBy string  `db:"complaint_by" json:"complaintby"`
}
