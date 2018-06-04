package api

import (
	"fmt"

	"opensoach.com/core/logger"
	"opensoach.com/hkt/constants"
	hktmodels "opensoach.com/hkt/models"
	"opensoach.com/hkt/server/dbaccess"
	repo "opensoach.com/hkt/server/repository"
	pcconst "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDeviceSPAssociated(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {
	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskSPDevAsscociatedModel := ctx.TaskData.(*hktmodels.TaskSPDevAsscociatedModel)

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskSPDevAsscociatedModel.CpmId)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching instance db connection.", dbErr)
		return dbErr, apiTaskProcessorResultModel
	}

	dbSerErr, deviceSerConfigDataList := dbaccess.TaskGetSerConfDetails(instDBConn,
		taskSPDevAsscociatedModel.CpmId, taskSPDevAsscociatedModel.DevId,
		taskSPDevAsscociatedModel.SpId)

	if dbSerErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching service configuration by servinstconfigid.", dbSerErr)
		return dbSerErr, apiTaskProcessorResultModel
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, deviceSerConfigData := range deviceSerConfigDataList {
		deviceTokenKey := fmt.Sprintf("%s%d", pcconst.CACHE_DEVICE_TOKEN_MAPPING_KEY_PREFIX, deviceSerConfigData.DeviceId)
		fmt.Println(deviceTokenKey)

		isTokenGetSucc, deviceToken := repo.Instance().Context.Master.Cache.Get(deviceTokenKey)

		if isTokenGetSucc == false {
			logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "ProcessDeviceSPAssociated:Unable to get device token. Device is offline. Skipping creation of packet")
		}

		epTaskSendPacketDataModel := pcmodels.EPTaskSendPacketDataModel{}
		epTaskSendPacketDataModel.Token = deviceToken
		epTaskSendPacketDataModel.Data = deviceSerConfigData
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_SERV_CONF

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel

	return nil, apiTaskProcessorResultModel
}

func ProcessSerConfigOnSP(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskSerConfigAddedOnSPModel := ctx.TaskData.(*hktmodels.TaskSerConfigAddedOnSPModel)

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskSerConfigAddedOnSPModel.CpmId)

	if dbErr != nil {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "ProcessSerConfigOnSP:Unable to get device token. Device is offline. Skipping creation of packet")
		return dbErr, apiTaskProcessorResultModel
	}

	dbSerErr, deviceSerConfigDataList := dbaccess.TaskGetSerConfDetailsByConfInstId(instDBConn, taskSerConfigAddedOnSPModel.ServInstConfID)

	if dbSerErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching service configuration by servinstconfigid.", dbSerErr)
		return dbSerErr, apiTaskProcessorResultModel
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, deviceSerConfigData := range deviceSerConfigDataList {
		deviceTokenKey := fmt.Sprintf("%s%d", pcconst.CACHE_DEVICE_TOKEN_MAPPING_KEY_PREFIX, deviceSerConfigData.DeviceId)
		fmt.Println(deviceTokenKey)

		isTokenGetSucc, deviceToken := repo.Instance().Context.Master.Cache.Get(deviceTokenKey)

		if isTokenGetSucc == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "", nil)
			continue
		}

		epTaskSendPacketDataModel := pcmodels.EPTaskSendPacketDataModel{}
		epTaskSendPacketDataModel.Token = deviceToken
		epTaskSendPacketDataModel.Data = deviceSerConfigData
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_SERV_CONF

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}

func ProcessSerConfigUpdated(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskServConfigUpdatedModel := ctx.TaskData.(*hktmodels.TaskServConfigUpdatedModel)

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskServConfigUpdatedModel.CpmId)

	if dbErr != nil {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "ProcessSerConfigEditedOnSP:Unable to get device token. Device is offline. Skipping creation of packet")
		return dbErr, apiTaskProcessorResultModel
	}

	dbSerErr, deviceSerConfigDataList := dbaccess.TaskGetSerConfDetailsByConfId(instDBConn, taskServConfigUpdatedModel.ServConfId)

	if dbSerErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching service configuration by servinstconfigid.", dbSerErr)
		return dbSerErr, apiTaskProcessorResultModel
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, deviceSerConfigData := range deviceSerConfigDataList {
		deviceTokenKey := fmt.Sprintf("%s%d", pcconst.CACHE_DEVICE_TOKEN_MAPPING_KEY_PREFIX, deviceSerConfigData.DeviceId)
		fmt.Println(deviceTokenKey)

		isTokenGetSucc, deviceToken := repo.Instance().Context.Master.Cache.Get(deviceTokenKey)

		if isTokenGetSucc == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "", nil)
			continue
		}

		epTaskSendPacketDataModel := pcmodels.EPTaskSendPacketDataModel{}
		epTaskSendPacketDataModel.Token = deviceToken
		epTaskSendPacketDataModel.Data = deviceSerConfigData
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_SERV_CONF

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}
