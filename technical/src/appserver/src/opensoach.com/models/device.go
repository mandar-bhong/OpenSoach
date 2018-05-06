package models

type DevicePacket struct {
	Header  DeviceHeaderData `json:"header"`
	Payload interface{}      `json:"payload"`
}

type DeviceHeaderData struct {
	CRC        string `json:"crc"`
	Category   int    `json:"category"`
	CommandID  int    `json:"commandid"`
	SeqID      int    `json:"seqid"`
	LocationID int    `json:"locationid"`
	Ack        int    `json:"ack"`
}
