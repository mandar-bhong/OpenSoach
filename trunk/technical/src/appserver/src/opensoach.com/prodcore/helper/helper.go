package helper

import (
	"strconv"

	gmodels "opensoach.com/models"
)

func GetDeviceCmdKeyFromHeader(deviceHeader gmodels.DeviceHeaderData) string {
	return strconv.Itoa(deviceHeader.Category) + "_" +
		strconv.Itoa(deviceHeader.CommandID) + "_" +
		strconv.Itoa(deviceHeader.Ack)
}

func GetDeviceCmdKey(category int, command int, ack int) string {
	return strconv.Itoa(category) + "_" +
		strconv.Itoa(command) + "_" +
		strconv.Itoa(ack)
}
