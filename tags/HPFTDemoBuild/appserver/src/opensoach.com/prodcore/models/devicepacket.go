package models

import (
	gmodels "opensoach.com/models"
)

type DevicePacketAuth struct {
	Token string `json:"token"`
}

type DevicePacketProccessExecution struct {
	Token          string
	DevicePacket   []byte
	InstanceDBConn string
	TokenInfo      *gmodels.DeviceTokenModel
}
