package processor

import (
	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	lhelper "opensoach.com/prodcore/helper"
	pcmodels "opensoach.com/prodcore/models"
)

type authSuccessHandler func(cacheCtx gcore.CacheContext, token string, chnID int)

func AuthorizeDevice(mstCache gcore.CacheContext, packet *gmodels.PacketProcessingTaskModel, successHandler authSuccessHandler) *gmodels.PacketProcessingTaskResult {

	result := &gmodels.PacketProcessingTaskResult{}

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

	//isSuccess, deviceAuthData := lhelper.CacheGetDeviceInfo(mstCache, payload.Token)
	isSuccess, _, _ := lhelper.CacheGetDeviceInfo(mstCache, payload.Token)

	if isSuccess == false {
		logger.Context().WithField("JSON Data", string(packet.Message)).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert from JSON packet", err)
		result.IsSuccess = false
		result.StatusCode = gmodels.MOD_OPER_ERR_SERVER
		return result
	}

	//Info: Handle local cache updation e.g. chnId vs token, token vs chnid
	successHandler(mstCache, payload.Token, packet.ChannelID)

	//INFO: Storing Entity ID vs connection status in cache
	//TODO: Store this information into cache
	//	isStatusSetSuccess := lhelper.CacheSetDeviceConnectionStatus(mstCache, deviceAuthData.DevID, true)

	//	if isStatusSetSuccess {
	//		logger.Context().WithField("CacheAddress", mstCache.CacheAddress).Log(SUB_MODULE_NAME, logger.Server, logger.Error, "Unable to set device connection status into cache")
	//		result.IsSuccess = false
	//		result.StatusCode = gmodels.MOD_OPER_ERR_SERVER
	//		return result
	//	}

	result.IsSuccess = true

	deviceAuthPacket := &gmodels.DevicePacket{}
	deviceAuthPacket.Header.Category = pcconst.DEVICE_CMD_CAT_DEVICE_VALIDATION
	deviceAuthPacket.Header.CommandID = pcconst.DEVICE_CMD_DEVICE_AUTH

	result.AckPayload = append(result.AckPayload, deviceAuthPacket)

	return result
}
