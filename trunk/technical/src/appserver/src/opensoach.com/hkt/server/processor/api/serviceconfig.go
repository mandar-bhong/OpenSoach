package api

import (
	"fmt"

	"opensoach.com/core/logger"
	hktmodels "opensoach.com/hkt/models"
	"opensoach.com/hkt/server/dbaccess"
	repo "opensoach.com/hkt/server/repository"
	pcconst "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDeviceSPAssociated(jsonmsg string) error {

	return nil
}

func ProcessSerConfigOnSP(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskSerConfigAddedOnSPModel := ctx.TaskData.(*hktmodels.TaskSerConfigAddedOnSPModel)

	fmt.Printf("Received taskSerConfigAddedOnSPModel : %#v \n",taskSerConfigAddedOnSPModel)

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskSerConfigAddedOnSPModel.CpmId)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching instance db connection.", dbErr)
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
			//TODO Log
		}

		epTaskSendPacketDataModel := pcmodels.EPTaskSendPacketDataModel{}
		epTaskSendPacketDataModel.Token = deviceToken
		epTaskSendPacketDataModel.Data = deviceSerConfigData
		epTaskSendPacketDataModel.TaskType = "ServiceConfig"

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}
