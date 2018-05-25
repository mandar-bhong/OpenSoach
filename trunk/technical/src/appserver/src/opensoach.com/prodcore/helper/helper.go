package helper

import (
	"strconv"

	gmodels "opensoach.com/models"
)

func GetDeviceCmdKeyFromHeader(deviceHeader gmodels.DeviceHeaderData) string {
	return strconv.Itoa(deviceHeader.Category) + "_" +
		strconv.Itoa(deviceHeader.CommandID)
}

func GetDeviceCmdKey(category int, command int) string {
	return strconv.Itoa(category) + "_" +
		strconv.Itoa(command)
}
