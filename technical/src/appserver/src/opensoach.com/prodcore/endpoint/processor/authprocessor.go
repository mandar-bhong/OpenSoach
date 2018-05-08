package processor

import (
	coremodels "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func AuthorizeDevice(mstCache coremodels.CacheContext, authMap map[int]string, packet *gmodels.PacketProcessingTaskModel) gmodels.PacketProcessingTaskResult {

	result := gmodels.PacketProcessingTaskResult{}

	payload := &pcmodels.DevicePacketAuth{}
	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = payload

	err := ghelper.ConvertFromJSONBytes(packet.Message, devicePacket)

	if err != nil {
		logger.Context().WithField("JSON Data", string(packet.Message)).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert from JSON packet", err)
		result.StatusCode = gmodels.MOD_OPER_ERR_SERVER
		result.IsSuccess = false
		return result
	}

	isSuccess, _ := mstCache.Get(payload.Token)

	if isSuccess == false {
		logger.Context().WithField("JSON Data", string(packet.Message)).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert from JSON packet", err)
		result.IsSuccess = false
		result.StatusCode = gmodels.MOD_OPER_ERR_SERVER
		return result
	}

	authMap[packet.ChannelID] = payload.Token

	result.IsSuccess = true

	deviceAuthPacket := &gmodels.DevicePacket{}
	deviceAuthPacket.Header.Category = 1
	deviceAuthPacket.Header.CommandID = 1
	deviceAuthPacket.Header.Ack = 1

	result.AckPayload = append(result.AckPayload, deviceAuthPacket)

	return result
}

func GetUnauthorizedDevicePacket() (bool, string) {

	devicePacket := gmodels.DevicePacket{}

	devicePacket.Header.Category = 1
	devicePacket.Header.CommandID = 1
	devicePacket.Header.Ack = 1

	return ghelper.ConvertToJSON(devicePacket)
}
