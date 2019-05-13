package helper

import (
	"strings"

	"opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	lconst "opensoach.com/hpft/server/constants"
	gmodels "opensoach.com/models"
	pcconstants "opensoach.com/prodcore/constants"
)

func GetEPTokenInfo(cache core.CacheContext, token string) (bool, int, *gmodels.DeviceTokenModel, *gmodels.DeviceUserSessionInfo) {

	isSuccess, data := cache.Get(token)
	if isSuccess == false {
		return false, pcconstants.DEVICE_TYPE_NONE, nil, nil
	}

	var contextType int
	tokenInfo := &gmodels.DeviceTokenModel{}
	userTokenInfo := &gmodels.DeviceUserSessionInfo{}

	if strings.HasPrefix(token, pcconstants.SHARED_DEVICE_TOKEN_PREFIX) {
		contextType = pcconstants.DEVICE_TYPE_SHARED_DEVICE
		if isJsonSuccess := ghelper.ConvertFromJSONString(data, tokenInfo); isJsonSuccess == false {
			return false, contextType, nil, nil
		}
	} else if strings.HasPrefix(token, pcconstants.USER_DEVICE_TOKEN_PREFIX) {
		contextType = pcconstants.DEVICE_TYPE_USER_DEVICE
		if isJsonSuccess := ghelper.ConvertFromJSONString(data, userTokenInfo); isJsonSuccess == false {
			return false, contextType, nil, nil
		}
	}

	return true, contextType, tokenInfo, userTokenInfo

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
