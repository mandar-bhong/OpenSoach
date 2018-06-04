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

func ProcessFieldOperatorOnSP(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskFieldOperatorAddedRemovedOnSPModel := ctx.TaskData.(*hktmodels.TaskFieldOperatorAddedRemovedOnSPModel)

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskFieldOperatorAddedRemovedOnSPModel.CpmId)

	if dbErr != nil {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "ProcessFieldOperatorOnSP:Unable to get device token. Device is offline. Skipping creation of packet")
		return dbErr, apiTaskProcessorResultModel
	}

	dbSerErr, deviceFieldOperatorDataList := dbaccess.TaskGetFieldOperatorDetailsByFopId(instDBConn, taskFieldOperatorAddedRemovedOnSPModel.FopId)

	if dbSerErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching field operator by fopid.", dbSerErr)
		return dbSerErr, apiTaskProcessorResultModel
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, deviceFieldOperatorData := range deviceFieldOperatorDataList {
		deviceTokenKey := fmt.Sprintf("%s%d", pcconst.CACHE_DEVICE_TOKEN_MAPPING_KEY_PREFIX, deviceFieldOperatorData.DeviceId)
		fmt.Println(deviceTokenKey)

		isTokenGetSucc, deviceToken := repo.Instance().Context.Master.Cache.Get(deviceTokenKey)

		if isTokenGetSucc == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "", nil)
			continue
		}

		epTaskSendPacketDataModel := pcmodels.EPTaskSendPacketDataModel{}
		epTaskSendPacketDataModel.Token = deviceToken
		epTaskSendPacketDataModel.Data = deviceFieldOperatorData
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_FIELD_OPERATOR

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}
