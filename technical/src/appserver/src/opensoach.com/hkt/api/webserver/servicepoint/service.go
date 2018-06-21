package servicepoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	repo "opensoach.com/hkt/api/repository"
	"opensoach.com/hkt/api/webserver/servicepoint/dbaccess"
	hktconst "opensoach.com/hkt/constants"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Servicepoint"

type ServicePointService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service ServicePointService) SpCategoryAdd(req lmodels.APISpCategoryAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBSpCategoryInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.SpcName = req.SpcName
	dbRowModel.ShortDesc = req.ShortDesc

	dbErr, insertedId := dbaccess.SpCategoryInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New SpCategory Added succesfully")

	return true, addResponse
}

func (service ServicePointService) SpUpdate(reqData *hktmodels.DBSpUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.SpStateSince = ghelper.GetCurrentTime()

	dbErr, affectedRow := dbaccess.SpUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Point updated successfully.")

	return true, nil
}

func (service ServicePointService) ServicePointAdd(req lmodels.APISpAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBSpInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.SpId = req.SpId
	dbRowModel.SpcId = req.SpcId
	dbRowModel.SpName = req.SpName
	dbRowModel.ShortDesc = req.ShortDesc
	dbRowModel.SpState = req.SpState
	dbRowModel.SpStateSince = ghelper.GetCurrentTime()

	dbErr, insertedId := dbaccess.SpInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New Service Point Added succesfully")

	return true, addResponse
}

func (service ServicePointService) SpCategoryShortDataList() (bool, interface{}) {

	cpmID := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, listData := dbaccess.GetSpCategoryShortDataList(service.ExeCtx.SessionInfo.Product.NodeDbConn, cpmID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Service Point Category short data list.")

	return true, listData

}

func (service ServicePointService) DevSpAssociation(req lmodels.APIDevSpAsscociationRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBDevSpMappingInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.DevId = req.DevId
	dbRowModel.SpId = req.SpId

	dbErr, insertedId := dbaccess.DevSpMappingTableInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	taskSPDevAsscociatedModel := &hktmodels.TaskSPDevAsscociatedModel{}
	taskSPDevAsscociatedModel.CpmId = dbRowModel.CpmId
	taskSPDevAsscociatedModel.DevId = dbRowModel.DevId
	taskSPDevAsscociatedModel.SpId = dbRowModel.SpId

	isSendSuccess := repo.Instance().
		SendTaskToServer(hktconst.TASK_HKT_API_DEVICE_SP_ASSOCIATED,
			service.ExeCtx.SessionToken, taskSPDevAsscociatedModel)

	if isSendSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device associated with Service Point  succesfully")

	return true, addResponse
}

func (service ServicePointService) DevSpAsscociationRemove(reqdata *lmodels.APIDevSpAsscociationRemoveRequest) (isSuccess bool, successErrorData interface{}) {

	dbErr, affectedRow := dbaccess.DevSpMappingTableDelete(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqdata)
	if dbErr != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device association with Service Point removed successfully.")

	return true, nil
}

func (service ServicePointService) GetSPList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchServicePointRequestFilterDataModel)

	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetServicePointList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched service point list data.")

	return true, dataListResponse

}

func (service ServicePointService) ServicePointShortDataList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetServicePointShortDataList(service.ExeCtx.SessionInfo.Product.NodeDbConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Service Point short data list.")

	return true, listData

}

func (service ServicePointService) GetServicePointInfo(spID int64) (bool, interface{}) {

	dbErr, spData := dbaccess.ServicePointSelectByID(service.ExeCtx.SessionInfo.Product.NodeDbConn, spID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *spData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched ServicePOint info")
	return true, dbRecord[0]
}

func (service ServicePointService) ServicePointConfigShortDataList() (bool, interface{}) {

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
