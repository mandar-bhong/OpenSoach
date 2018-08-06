package api

import (
	"fmt"

	"opensoach.com/core/logger"
	"opensoach.com/hpft/constants"
	hktmodels "opensoach.com/hpft/models"
	"opensoach.com/hpft/server/dbaccess"
	repo "opensoach.com/hpft/server/repository"
	pcconst "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessPatientStatusUpdated(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskPatientStatusUpdated := ctx.TaskData.(*hktmodels.TaskPatientStatusUpdated)

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskPatientStatusUpdated.CpmId)

	if dbErr != nil {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching instance db connection")
		return dbErr, apiTaskProcessorResultModel
	}

	//get devices by service points
	dbErr, devspdata := dbaccess.TaskGetDeviceByPatientID(instDBConn, taskPatientStatusUpdated.PatientId)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching device by patient id.", dbErr)
		return dbErr, apiTaskProcessorResultModel
	}

	dbSerErr, devicePatientConfigDataList := dbaccess.TaskGetPatientConfDetails(instDBConn, taskPatientStatusUpdated.PatientId)

	fmt.Println(devicePatientConfigDataList)

	if dbSerErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching service configuration by servinstconfigid.", dbSerErr)
		return dbSerErr, apiTaskProcessorResultModel
	}

	dbDevicePatientConfigModelList := []hktmodels.DBDevicePatientConfigModel{}

	for i := 0; i < len(devspdata); i++ {
		for j := 0; i < len(devicePatientConfigDataList); i++ {
			dbDevicePatientConfigModel := hktmodels.DBDevicePatientConfigModel{}
			dbDevicePatientConfigModel.SpId = devspdata[i].SpId
			dbDevicePatientConfigModel.DeviceId = devspdata[i].DevId
			dbDevicePatientConfigModel.SerConfId = devicePatientConfigDataList[j].SerConfId
			dbDevicePatientConfigModel.SerConfInstId = devicePatientConfigDataList[j].SerConfInstId
			dbDevicePatientConfigModel.ServConfCode = devicePatientConfigDataList[j].ServConfCode
			dbDevicePatientConfigModel.ServConfName = devicePatientConfigDataList[j].ServConfName
			dbDevicePatientConfigModel.ServiceConfig = devicePatientConfigDataList[j].ServiceConfig
			dbDevicePatientConfigModel.PatientDetails = devicePatientConfigDataList[j].PatientDetails
			dbDevicePatientConfigModel.MedicalDetails = devicePatientConfigDataList[j].MedicalDetails
			dbDevicePatientConfigModelList = append(dbDevicePatientConfigModelList, dbDevicePatientConfigModel)
		}
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, dbDevicePatientConfigdata := range dbDevicePatientConfigModelList {
		deviceTokenKey := fmt.Sprintf("%s%d", pcconst.CACHE_DEVICE_TOKEN_MAPPING_KEY_PREFIX, dbDevicePatientConfigdata.DeviceId)
		fmt.Println(deviceTokenKey)

		isTokenGetSucc, deviceToken := repo.Instance().Context.Master.Cache.Get(deviceTokenKey)

		if isTokenGetSucc == false {
			logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "ProcessPatientStatusUpdated:Unable to get device token. Device is offline. Skipping creation of packet")
			continue
		}

		epTaskSendPacketDataModel := pcmodels.EPTaskSendPacketDataModel{}
		epTaskSendPacketDataModel.Token = deviceToken
		epTaskSendPacketDataModel.Data = dbDevicePatientConfigdata
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_PATIENT_CONF

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}
