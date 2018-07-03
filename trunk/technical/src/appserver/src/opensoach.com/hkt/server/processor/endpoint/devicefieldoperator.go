package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	"opensoach.com/hkt/constants"
	hktmodels "opensoach.com/hkt/models"
	lconst "opensoach.com/hkt/server/constants"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDeviceFieldOperator(epTaskSendPacketDataModelList []pcmodels.EPTaskSendPacketDataModel) {

	var tokenSPServices map[string]map[int64][]hktmodels.DBDeviceFieldOperatorDataModel
	tokenSPServices = make(map[string]map[int64][]hktmodels.DBDeviceFieldOperatorDataModel)

	var CommandID int

	for _, epTask := range epTaskSendPacketDataModelList {

		switch epTask.TaskType {

		case constants.TASK_TYPE_FIELD_OPERATOR_ASSOCIATED:

			CommandID = lconst.DEVICE_CMD_CONFIG_SERVICE_POINTS_FIELD_OPERATOR_ASSOCIATED

		case constants.TASK_TYPE_FIELD_OPERATOR_DEASSOCIATED:

			CommandID = lconst.DEVICE_CMD_CONFIG_SERVICE_POINTS_FIELD_OPERATOR_DEASSOCIATED

		case constants.TASK_TYPE_FIELD_OPERATOR_ADDED:

			CommandID = lconst.DEVICE_CMD_CONFIG_SERVICE_POINTS_FIELD_OPERATOR_ADDED

		}

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

			isJsonSucc, jsonData := createFieldOperatorPacket(spid, spFopItem, CommandID)

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

func createFieldOperatorPacket(spid int64, fopModels []hktmodels.DBDeviceFieldOperatorDataModel, commandid int) (bool, string) {

	fopinfo := &gmodels.DevicePacket{}
	fopinfo.Header = gmodels.DeviceHeaderData{}
	fopinfo.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
	fopinfo.Header.CommandID = commandid
	fopinfo.Header.SPID = spid

	fopList := []string{}

	for _, fopModel := range fopModels {
		fopList = append(fopList, fopModel.Fopcode)
	}

	fopinfo.Payload = fopList

	isJsonSucc, jsonData := ghelper.ConvertToJSON(fopinfo)

	return isJsonSucc, jsonData

}
