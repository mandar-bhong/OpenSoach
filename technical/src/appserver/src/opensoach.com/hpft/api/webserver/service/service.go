package service

import (
	"time"

	"opensoach.com/core/logger"
	lmodels "opensoach.com/hpft/api/models"
	repo "opensoach.com/hpft/api/repository"
	"opensoach.com/hpft/api/webserver/service/dbaccess"
	hktconst "opensoach.com/hpft/constants"
	hktmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.Service"

type ServiceConfigService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service ServiceConfigService) ServiceConfigAdd(req lmodels.APIServiceConfAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBServiceConfInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.SpcId = req.SpcId
	dbRowModel.ConfTypeCode = req.ConfTypeCode
	dbRowModel.ServConfName = req.ServConfName
	dbRowModel.ShortDesc = req.ShortDesc
	dbRowModel.ServConf = req.ServConf

	dbErr, insertedId := dbaccess.ServiceConfigInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New ServiceConf Added succesfully")

	return true, addResponse
}

func (service ServiceConfigService) ServiceConfigList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchServiceConfRequestFilterModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetServiceConfigList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Service Conf list data.")

	return true, dataListResponse

}

func (service ServiceConfigService) ServiceConnfigUpdate(reqData *hktmodels.DBServiceConfUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, affectedRow := dbaccess.ServiceConfigUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	taskServConfigUpdatedModel := &hktmodels.TaskServConfigUpdatedModel{}
	taskServConfigUpdatedModel.ServConfId = reqData.ServConfId
	taskServConfigUpdatedModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	isSendSuccess := repo.Instance().
		SendTaskToServer(hktconst.TASK_HKT_API_SERVICE_CONFIG_UPDATED,
			service.ExeCtx.SessionToken, taskServConfigUpdatedModel)

	if isSendSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Conf data updated successfully.")

	return true, nil
}

func (service ServiceConfigService) ServiceInstanceAdd(req lmodels.APIServiceInstanceAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBServiceInstanceInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.ServConfId = req.ServConfId
	dbRowModel.SpId = req.SpId

	dbErr, insertedId := dbaccess.ServiceInstanceInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding service instance data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	taskSerConfigAddedOnSPModel := &hktmodels.TaskSerConfigAddedOnSPModel{}
	taskSerConfigAddedOnSPModel.ServInstConfID = insertedId
	taskSerConfigAddedOnSPModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	isSendSuccess := repo.Instance().
		SendTaskToServer(hktconst.TASK_HKT_API_SERVICE_CONFIG_ADDED_ON_SP,
			service.ExeCtx.SessionToken, taskSerConfigAddedOnSPModel)

	if isSendSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New ServiceInstance Added succesfully")

	return true, addResponse
}

func (service ServiceConfigService) ServiceinstanceList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchServiceInstanceRequestFilterModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetServiceInstanceList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched serviceconf list data.")

	return true, dataListResponse

}

func (service ServiceConfigService) GetServiceInstanceTxn(spid int, StartDate time.Time, EndDate time.Time) (bool, interface{}) {

	cpmID := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, complaintList := dbaccess.GetServiceInstTxn(service.ExeCtx.SessionInfo.Product.NodeDbConn, cpmID, spid, StartDate, EndDate)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched service instance  transaction data")
	return true, complaintList
}

func (service ServiceConfigService) ServiceConfShortDataList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetServiceConfShortDataList(service.ExeCtx.SessionInfo.Product.NodeDbConn, service.ExeCtx.SessionInfo.Product.CustProdID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Service Conf short data list.")

	return true, listData

}

func (service ServiceConfigService) ServiceConfigCopyTemplate(req hktmodels.DBServiceConfTemplateInsertDataModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, insertedId := dbaccess.ServiceConfigInsertCopy(service.ExeCtx.SessionInfo.Product.NodeDbConn, req)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New Service Config Template copied succesfully")

	return true, addResponse
}

func (service ServiceConfigService) ServiceConfInfo(servconfID int64) (bool, interface{}) {

	dbErr, servConfData := dbaccess.ServiceConfSelectByID(service.ExeCtx.SessionInfo.Product.NodeDbConn, servconfID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *servConfData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Service Conf info")
	return true, dbRecord[0]
}

func (service ServiceConfigService) ServicePointConfigShortDataList() (bool, interface{}) {

	cpmID := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, listData := dbaccess.GetServicePointConfigList(service.ExeCtx.SessionInfo.Product.NodeDbConn, cpmID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Service Point Config short data list.")

	return true, listData

}
