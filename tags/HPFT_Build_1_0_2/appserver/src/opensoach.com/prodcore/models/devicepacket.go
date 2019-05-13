package models

type DevicePacketAuth struct {
	Token string `json:"token"`
}

type DevicePacketProccessExecution struct {
	Token          string
	DevicePacket   []byte
	InstanceDBConn string
	DeviceContext  interface{}
}
