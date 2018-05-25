package helper

import (
	"strconv"

	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
)

func GetDeviceCmdKeyFromHeader(deviceHeader gmodels.DeviceHeaderData) string {
	return strconv.Itoa(deviceHeader.Category) + "_" +
		strconv.Itoa(deviceHeader.CommandID)
}

func GetDeviceCmdKey(category int, command int) string {
	return strconv.Itoa(category) + "_" +
		strconv.Itoa(command)
}

func GetEPAckPacket(commandID int, seqid int, isSuccess bool, errorCode int, ackData interface{}) *gmodels.DevicePacket {
	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Header.Category = pcconst.DEVICE_CMD_CAT_ACK
	devicePacket.Header.CommandID = commandID
	devicePacket.Header.SeqID = seqid

	deviceCommandAck := gmodels.DeviceCommandAck{}
	deviceCommandAck.Ack = isSuccess
	deviceCommandAck.Data = ackData
	devicePacket.Payload = deviceCommandAck

	return devicePacket
}
