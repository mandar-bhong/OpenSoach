package processor

import (
	coremodels "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func AuthorizeDevice(mstCache coremodels.CacheContext, authMap map[int]string, packet *gmodels.EndPointToServerTaskModel) gmodels.PacketProcessingResult {

	result := gmodels.PacketProcessingResult{}

	payload := &pcmodels.DevicePacketAuth{}
	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = payload

	err := ghelper.ConvertFromJSONBytes(packet.Message, devicePacket)

	if err != nil {
		//Error
		return result
	}

	isSuccess, _ := mstCache.Get(payload.Token)

	//packet.AuthData = authData

	if isSuccess == false {
		//Error
		return result
	}

	authMap[packet.ChannelID] = payload.Token

	result.IsSuccess = true

	return result
}
