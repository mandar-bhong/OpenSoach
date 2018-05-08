package models

import (
	gmodels "opensoach.com/models"
)

type PacketProccessExecution struct {
	Token          string
	DevicePacket   []byte
	InstanceDBConn string
	TokenInfo      *gmodels.DeviceTokenModel
}
