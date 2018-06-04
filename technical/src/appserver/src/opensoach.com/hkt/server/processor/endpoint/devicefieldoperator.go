package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	hktmodels "opensoach.com/hkt/models"
	lconst "opensoach.com/hkt/server/constants"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDeviceFieldOperator(epTaskSendPacketDataModelList []pcmodels.EPTaskSendPacketDataModel) {

	var tokenSPServices map[string]map[int64][]hktmodels.DBDeviceFieldOperatorDataModel
	tokenSPServices = make(map[string]map[int64][]hktmodels.DBDeviceFieldOperatorDataModel)

	for _, epTask := range epTaskSendPacketDataModelList {

		_, hasToken := tokenSPServices[epTask.Token]

		if hasToken == false {
			tokenSPServices[epTask.Token] = map[int64][]hktmodels.DBDeviceFieldOperatorDataModel{}
		}

		dbDeviceFieldOperatorDataModel := epTask.Data.(hktmodels.DBDeviceFieldOperatorDataModel)

		_, hasSP := tokenSPServices[epTask.Token][dbDeviceFieldOperatorDataModel.SpId]

		if hasSP == false {
			tokenSPServices[epTask.Token][dbDeviceFieldOperatorDataModel.SpId] = []hktmodels.DBDeviceFieldOperatorDataModel{dbDeviceFieldOperatorDataModel}
		} else {
			tokenSPServices[epTask.Token][dbDeviceFieldOperatorDataModel.SpId] = append(tokenSPServices[epTask.Token][dbDeviceFieldOperatorDataModel.SpId], dbDeviceFieldOperatorDataModel)
		}
	}

	epPackets := []pcmodels.TaskEPPacketSendDataModel{}

	for token, tokenSPItem := range tokenSPServices {

		for spid, spFopItem := range tokenSPItem {

			isJsonSucc, jsonData := createFieldOperatorPacket(spid, spFopItem)

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

func createFieldOperatorPacket(spid int64, fopModels []hktmodels.DBDeviceFieldOperatorDataModel) (bool, string) {

	fopinfo := &gmodels.DevicePacket{}
	fopinfo.Header = gmodels.DeviceHeaderData{}
	fopinfo.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
	fopinfo.Header.CommandID = lconst.DEVICE_CMD_CONFIG_SERVICE_POINTS_FIELD_OPERATOR
	fopinfo.Header.SPID = spid

	var fopList []hktmodels.DBEPSPFieldOperatorDataModel

	for _, fopModel := range fopModels {
		dbEPSPFieldOperatorDataModel := hktmodels.DBEPSPFieldOperatorDataModel{}

		dbEPSPFieldOperatorDataModel.FopId = fopModel.FopId
		dbEPSPFieldOperatorDataModel.Fopcode = fopModel.Fopcode
		dbEPSPFieldOperatorDataModel.FopName = fopModel.FopName
		dbEPSPFieldOperatorDataModel.FopState = fopModel.FopState
		dbEPSPFieldOperatorDataModel.FopArea = fopModel.FopArea

		fopList = append(fopList, dbEPSPFieldOperatorDataModel)
	}

	fopinfo.Payload = fopList

	isJsonSucc, jsonData := ghelper.ConvertToJSON(fopinfo)

	return isJsonSucc, jsonData

}
