package api

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	"opensoach.com/hpft/constants"
	hktmodels "opensoach.com/hpft/models"
	"opensoach.com/hpft/server/dbaccess"
	repo "opensoach.com/hpft/server/repository"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	pchelper "opensoach.com/prodcore/helper"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessFieldOperatorAddedOnSP(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

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
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_FIELD_OPERATOR_ASSOCIATED

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}

func ProcessFieldOperatorRemovedOnSP(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskFieldOperatorAddedRemovedOnSPModel := ctx.TaskData.(*hktmodels.TaskFieldOperatorAddedRemovedOnSPModel)

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskFieldOperatorAddedRemovedOnSPModel.CpmId)

	if dbErr != nil {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "ProcessFieldOperatorOnSP:Unable to get device token. Device is offline. Skipping creation of packet")
		return dbErr, apiTaskProcessorResultModel
	}

	dbSerErr, deviceFieldOperatorData := dbaccess.TaskGetFieldOperatorByFopId(instDBConn, taskFieldOperatorAddedRemovedOnSPModel.FopId)

	if dbSerErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching field operator by fopid.", dbSerErr)
		return dbSerErr, apiTaskProcessorResultModel
	}

	//get devices by service points
	dbErr, devspdata := dbaccess.TaskGetDeviceBySpID(instDBConn, taskFieldOperatorAddedRemovedOnSPModel.SpId)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching field operator by fopid.", dbErr)
		return dbErr, apiTaskProcessorResultModel
	}

	dbDeviceFieldOperatorDataModelList := []hktmodels.DBDeviceFieldOperatorDataModel{}

	for i := 0; i < len(devspdata); i++ {
		dbDeviceFieldOperatorDataModel := hktmodels.DBDeviceFieldOperatorDataModel{}
		dbDeviceFieldOperatorDataModel.Fopcode = deviceFieldOperatorData.Fopcode
		dbDeviceFieldOperatorDataModel.SpId = devspdata[i].SpId
		dbDeviceFieldOperatorDataModel.DeviceId = devspdata[i].DevId
		dbDeviceFieldOperatorDataModelList = append(dbDeviceFieldOperatorDataModelList, dbDeviceFieldOperatorDataModel)
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, deviceFieldOperatorData := range dbDeviceFieldOperatorDataModelList {
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
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_FIELD_OPERATOR_DEASSOCIATED

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}

func ProcessFieldOperatorAdded(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskFieldOperatorAddedModel := ctx.TaskData.(*hktmodels.TaskFieldOperatorAddedRemovedOnSPModel)

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskFieldOperatorAddedModel.CpmId)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "ProcessFieldOperatorOnSP:Unable to get device token. Device is offline. Skipping creation of packet", dbErr)
		return dbErr, apiTaskProcessorResultModel
	}

	//get online devices
	getErr, tokenlistjsonstring := repo.Instance().ProdTaskContext.ProcessTask(pcconst.TASK_GET_ONLINE_DEVICES, "")
	if getErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while submitting task.", getErr)
		return fmt.Errorf("Error occured while submitting task."), apiTaskProcessorResultModel
	}

	issuccess, deviceDataList := getOnlineDevices(tokenlistjsonstring)
	if issuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while submitting task.", nil)
		return fmt.Errorf("Error occured while fetching online devices."), apiTaskProcessorResultModel
	}

	//get device service points
	dbDeviceServicePointDataModelList := []hktmodels.DBDeviceServicePointDataModel{}

	for _, deviceData := range deviceDataList {

		dbErr, devspdata := dbaccess.TaskGetServicePointByDevId(instDBConn, deviceData.DevID)

		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching field operator by fopid.", dbErr)
			return dbErr, apiTaskProcessorResultModel
		}

		dbDeviceServicePointDataModelList = append(dbDeviceServicePointDataModelList, devspdata)

	}

	//get field operator data
	dbErr, deviceFieldOperatorData := dbaccess.TaskGetFieldOperatorByFopId(instDBConn, taskFieldOperatorAddedModel.FopId)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching field operator by fopid.", dbErr)
		return dbErr, apiTaskProcessorResultModel
	}

	dbDeviceFieldOperatorDataModelList := []hktmodels.DBDeviceFieldOperatorDataModel{}

	for i := 0; i < len(dbDeviceServicePointDataModelList); i++ {
		dbDeviceFieldOperatorDataModel := hktmodels.DBDeviceFieldOperatorDataModel{}
		dbDeviceFieldOperatorDataModel.Fopcode = deviceFieldOperatorData.Fopcode
		dbDeviceFieldOperatorDataModel.SpId = dbDeviceServicePointDataModelList[i].SpId
		dbDeviceFieldOperatorDataModel.DeviceId = dbDeviceServicePointDataModelList[i].DevId
		dbDeviceFieldOperatorDataModelList = append(dbDeviceFieldOperatorDataModelList, dbDeviceFieldOperatorDataModel)
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, deviceFieldOperatorData := range dbDeviceFieldOperatorDataModelList {
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
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_FIELD_OPERATOR_ADDED

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	apiTaskProcessorResultModel.EPSyncData = epTaskSendPacketDataList
	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}

func getOnlineDevices(tokenlistjsonstring string) (bool, []gmodels.DeviceTokenModel) {

	tokens := []string{}

	isJsonSuccess := ghelper.ConvertFromJSONString(tokenlistjsonstring, &tokens)

	if isJsonSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert json device packet ", nil)
		return false, nil
	}

	deviceTokenModelList := []gmodels.DeviceTokenModel{}

	for _, token := range tokens {

		isSuccess, _, jsonData := pchelper.CacheGetDeviceInfo(repo.Instance().Context.Master.Cache, token)

		if isSuccess == false {
			logger.Context().WithField("Token", token).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to get information for provided token")
			return false, nil
		}

		deviceTokenModel := gmodels.DeviceTokenModel{}
		isJsonSuccess := ghelper.ConvertFromJSONString(jsonData, &deviceTokenModel)

		if isJsonSuccess == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert json device packet ", nil)
			return false, nil
		}

		deviceTokenModelList = append(deviceTokenModelList, deviceTokenModel)
	}

	return true, deviceTokenModelList

}
