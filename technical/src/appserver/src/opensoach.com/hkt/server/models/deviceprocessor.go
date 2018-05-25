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
	TxnData           string    `json:"txndata"`
	TxnDate           time.Time `json:"txndate"`
	//	Status            int       `json:"status"`
	//	FOPCode           string    `json:"fopcode"`
}
