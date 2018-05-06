package processor

import (
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
)

func DecodeHeader(packet []byte) (error, gmodels.DeviceHeaderData) {
	deveicePacket := &gmodels.DevicePacket{}
	err := ghelper.ConvertFromJSONBytes(packet, deveicePacket)
	return err, deveicePacket.Header
}

//func GetCmdKeyFromHeader(deviceHeader gmodels.DeviceHeaderData) string {
//	return strconv.Itoa(deviceHeader.Category) + "_" +
//		strconv.Itoa(deviceHeader.CommandID) + "_" +
//		strconv.Itoa(deviceHeader.Ack)
//}

//func GetCmdKey(category int, command int, ack int) string {
//	return strconv.Itoa(category) + "_" +
//		strconv.Itoa(command) + "_" +
//		strconv.Itoa(ack)
//}
