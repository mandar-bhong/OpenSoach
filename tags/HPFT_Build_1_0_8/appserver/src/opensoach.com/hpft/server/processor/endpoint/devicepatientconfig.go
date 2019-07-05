package endpoint

import (
	"opensoach.com/core/logger"

	ghelper "opensoach.com/core/helper"
	hktmodels "opensoach.com/hpft/models"
	lconst "opensoach.com/hpft/server/constants"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDevicePatientConfig(epTaskSendPacketDataModelList []pcmodels.EPTaskSendPacketDataModel) {

	var tokenSPServices map[string]map[int64][]hktmodels.DBDevicePatientConfigModel
	tokenSPServices = make(map[string]map[int64][]hktmodels.DBDevicePatientConfigModel)

	for _, epTask := range epTaskSendPacketDataModelList {

		_, hasToken := tokenSPServices[epTask.Token]

		if hasToken == false {
			tokenSPServices[epTask.Token] = map[int64][]hktmodels.DBDevicePatientConfigModel{}
		}

		dbDevicePatientConfigModel := epTask.Data.(hktmodels.DBDevicePatientConfigModel)

		_, hasSP := tokenSPServices[epTask.Token][dbDevicePatientConfigModel.SpId]

		if hasSP == false {
			tokenSPServices[epTask.Token][dbDevicePatientConfigModel.SpId] = []hktmodels.DBDevicePatientConfigModel{dbDevicePatientConfigModel}
		} else {
			tokenSPServices[epTask.Token][dbDevicePatientConfigModel.SpId] = append(tokenSPServices[epTask.Token][dbDevicePatientConfigModel.SpId], dbDevicePatientConfigModel)
		}
	}

	epPackets := []pcmodels.TaskEPPacketSendDataModel{}

	for token, tokenSPItem := range tokenSPServices {

		for spid, spConfItem := range tokenSPItem {

			isJsonSucc, jsonData := createPatientConfigPacket(spid, spConfItem)

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

func createPatientConfigPacket(spid int64, serConfigModels []hktmodels.DBDevicePatientConfigModel) (bool, string) {

	servconfinfo := &gmodels.DevicePacket{}
	servconfinfo.Header = gmodels.DeviceHeaderData{}
	servconfinfo.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
	servconfinfo.Header.CommandID = lconst.DEVCIE_CMD_CONFIG_SERVICE_POINTS_PATIENT_CONF
	servconfinfo.Header.SPID = spid

	var serConfigList []hktmodels.DBEPSPPatientConfDataModel

	for _, serConfigModel := range serConfigModels {
		dbEPSPPatientConfDataModel := hktmodels.DBEPSPPatientConfDataModel{}

		dbEPSPPatientConfDataModel.ConfTypeCode = serConfigModel.ServConfCode
		dbEPSPPatientConfDataModel.ServConf = serConfigModel.ServiceConfig
		dbEPSPPatientConfDataModel.ServConfId = serConfigModel.SerConfId
		dbEPSPPatientConfDataModel.ServConfName = serConfigModel.ServConfName
		dbEPSPPatientConfDataModel.ServInId = serConfigModel.SerConfInstId
		dbEPSPPatientConfDataModel.PatientDetails = serConfigModel.PatientDetails
		dbEPSPPatientConfDataModel.MedicalDetails = serConfigModel.MedicalDetails

		serConfigList = append(serConfigList, dbEPSPPatientConfDataModel)
	}

	servconfinfo.Payload = serConfigList

	isJsonSucc, jsonData := ghelper.ConvertToJSON(servconfinfo)

	return isJsonSucc, jsonData

}
