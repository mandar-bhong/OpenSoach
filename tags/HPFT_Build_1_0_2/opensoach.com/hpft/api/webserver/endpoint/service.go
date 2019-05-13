package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	apimodels "opensoach.com/hpft/api/models"
	"opensoach.com/hpft/api/webserver/endpoint/dbaccess"
	hpftmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.Endpoint"

type EndpointService struct {
	ExeCtx *gmodels.DeviceUserExecutionContext
}

func (service EndpointService) GetPatientAdmissionList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	usrID := service.ExeCtx.DeviceUserSessionInfo.UserID
	filterModel := listReqData.Filter.(*hpftmodels.DBDeviceSearchPatientRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.DeviceUserSessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientList(service.ExeCtx.DeviceUserSessionInfo.Product.NodeDbConn, usrID, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient admission list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient admission list data.")

	return true, dataListResponse

}

func (service EndpointService) UserPateintAssociate(req apimodels.APIUserPatientAsscociationRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hpftmodels.DBPatientMonitorMappingInsertRowModel{}
	dbRowModel.DBPatientMonitorMappingDataModel = req.DBPatientMonitorMappingDataModel
	dbRowModel.CpmId = service.ExeCtx.DeviceUserSessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.DeviceUserSessionInfo.UserID
	dbRowModel.Uuid = ghelper.GenerateUUID()

	if req.PatientId == nil {
		dbPatientMonitorMappingDeleteRowModel := &hpftmodels.DBPatientMonitorMappingDeleteRowModel{}
		dbPatientMonitorMappingDeleteRowModel.CpmId = dbRowModel.CpmId
		dbPatientMonitorMappingDeleteRowModel.UsrId = dbRowModel.UsrId
		dbPatientMonitorMappingDeleteRowModel.SpId = dbRowModel.SpId
		service.UserPatientAsscociationRemove(dbPatientMonitorMappingDeleteRowModel)
	} else if req.SpId == nil {
		dbPatientMonitorMappingDeleteRowModel := &hpftmodels.DBPatientMonitorMappingDeleteRowModel{}
		dbPatientMonitorMappingDeleteRowModel.CpmId = dbRowModel.CpmId
		dbPatientMonitorMappingDeleteRowModel.UsrId = dbRowModel.UsrId
		service.UserPatientAsscociationRemove(dbPatientMonitorMappingDeleteRowModel)
	}

	dbErr, insertedId := dbaccess.PatientUserAssociation(service.ExeCtx.DeviceUserSessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while associated user and patient.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New user and patient associated succesfully")

	return true, addResponse
}

func (service EndpointService) UserPatientAsscociationRemove(reqdata *hpftmodels.DBPatientMonitorMappingDeleteRowModel) (isSuccess bool, successErrorData interface{}) {

	reqdata.CpmId = service.ExeCtx.DeviceUserSessionInfo.Product.CustProdID

	dbErr, affectedRow := dbaccess.PatientUserDeAssociation(service.ExeCtx.DeviceUserSessionInfo.Product.NodeDbConn, reqdata)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while deassociating user patient.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User patient association removed successfully.")

	return true, nil
}
