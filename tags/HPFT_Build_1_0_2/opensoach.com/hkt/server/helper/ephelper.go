package helper

import (
	"opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	lconst "opensoach.com/hkt/server/constants"
	gmodels "opensoach.com/models"
)

func GetEPTokenInfo(cache core.CacheContext, token string) (bool, *gmodels.DeviceTokenModel) {

	isSuccess, data := cache.Get(token)

	if isSuccess == false {
		return false, nil
	}
	tokenInfo := &gmodels.DeviceTokenModel{}

	if isJsonSuccess := ghelper.ConvertFromJSONString(data, tokenInfo); isJsonSuccess == false {
		return false, nil
	}

	return true, tokenInfo

}

func GetEPAckPacket(commandID int, seqid int, isSuccess bool, errorCode int, ackData interface{}) *gmodels.DevicePacket {
	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Header.Category = lconst.DEVICE_CMD_CAT_ACK
	devicePacket.Header.CommandID = commandID
	devicePacket.Header.SeqID = seqid

	deviceCommandAck := gmodels.DeviceCommandAck{}
	deviceCommandAck.Ack = isSuccess
	deviceCommandAck.Data = ackData
	devicePacket.Payload = deviceCommandAck

	return devicePacket
}
