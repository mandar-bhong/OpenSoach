package endpoint

import (
	"opensoach.com/core/logger"

	ghelper "opensoach.com/core/helper"
	hktmodels "opensoach.com/hpft/models"
	lconst "opensoach.com/hpft/server/constants"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDeviceServiceConfig(epTaskSendPacketDataModelList []pcmodels.EPTaskSendPacketDataModel) {

	var tokenSPServices map[string]map[int64][]hktmodels.DBDeviceSerConfigModel
	tokenSPServices = make(map[string]map[int64][]hktmodels.DBDeviceSerConfigModel)

	for _, epTask := range epTaskSendPacketDataModelList {

		_, hasToken := tokenSPServices[epTask.Token]

		if hasToken == false {
			tokenSPServices[epTask.Token] = map[int64][]hktmodels.DBDeviceSerConfigModel{}
		}

		dbDeviceSerConfigModel := epTask.Data.(hktmodels.DBDeviceSerConfigModel)

		_, hasSP := tokenSPServices[epTask.Token][dbDeviceSerConfigModel.SpId]

		if hasSP == false {
			tokenSPServices[epTask.Token][dbDeviceSerConfigModel.SpId] = []hktmodels.DBDeviceSerConfigModel{dbDeviceSerConfigModel}
		} else {
			tokenSPServices[epTask.Token][dbDeviceSerConfigModel.SpId] = append(tokenSPServices[epTask.Token][dbDeviceSerConfigModel.SpId], dbDeviceSerConfigModel)
		}
	}

	epPackets := []pcmodels.TaskEPPacketSendDataModel{}

	for token, tokenSPItem := range tokenSPServices {

		for spid, spConfItem := range tokenSPItem {

			isJsonSucc, jsonData := createServiceConfigPacket(spid, spConfItem)

			if isJsonSucc == false {
				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert to json data", nil)
				continue
			}

			packetSendDataModel := pcmodels.TaskEPPacketSendDataModel{}
			packetSendDataModel.Token = token
			packetSendDataModel.Packet = jsonData

			epPackets = append(epPackets, packetSendDataModel)
		}
	}

	SendPacketToEP(epPackets)
}

func createServiceConfigPacket(spid int64, serConfigModels []hktmodels.DBDeviceSerConfigModel) (bool, string) {

	servconfinfo := &gmodels.DevicePacket{}
	servconfinfo.Header = gmodels.DeviceHeaderData{}
	servconfinfo.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
	servconfinfo.Header.CommandID = lconst.DEVCIE_CMD_CONFIG_SERVICE_POINTS_SERV_CONF
	servconfinfo.Header.SPID = spid

	var serConfigList []hktmodels.DBEPSPServConfDataModel

	for _, serConfigModel := range serConfigModels {
		dbEPSPServConfDataModel := hktmodels.DBEPSPServConfDataModel{}

		dbEPSPServConfDataModel.ConfTypeCode = serConfigModel.ServConfCode
		dbEPSPServConfDataModel.ServConf = serConfigModel.ServiceConfig
		dbEPSPServConfDataModel.ServConfId = serConfigModel.SerConfId
		dbEPSPServConfDataModel.ServConfName = serConfigModel.ServConfName
		dbEPSPServConfDataModel.ServInId = serConfigModel.SerConfInstId

		serConfigList = append(serConfigList, dbEPSPServConfDataModel)
	}

	servconfinfo.Payload = serConfigList

	isJsonSucc, jsonData := ghelper.ConvertToJSON(servconfinfo)

	return isJsonSucc, jsonData

}
