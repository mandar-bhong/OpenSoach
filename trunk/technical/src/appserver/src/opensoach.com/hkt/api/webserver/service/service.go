package service

import (
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/service/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Service"

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Conf data updated successfully.")

	return true, nil
}

func (service ServiceConfigService) ServiceInstanceAdd(req lmodels.APIServiceInstanceAdddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBServiceInstanceInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.ServConfId = req.ServConfId
	dbRowModel.SpId = req.SpId

	dbErr, insertedId := dbaccess.ServiceInstanceInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New ServiceInstance Added succesfully")

	return true, addResponse
}

func (service ServiceConfigService) ServiceinstanceList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchServiceInstanceRequestFilterModel)

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
