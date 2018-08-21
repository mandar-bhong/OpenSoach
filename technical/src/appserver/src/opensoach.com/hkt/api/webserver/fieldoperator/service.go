package fieldoperator

import (
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	repo "opensoach.com/hkt/api/repository"
	"opensoach.com/hkt/api/webserver/fieldoperator/dbaccess"
	hktconst "opensoach.com/hkt/constants"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Fieldoperator"

type FieldoperatorService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service FieldoperatorService) Add(req lmodels.APIFieldOperatorAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbFieldOperatorRowModel := &hktmodels.DBFieldOperatorRowModel{}
	dbFieldOperatorRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbFieldOperatorRowModel.Fopcode = req.Fopcode
	dbFieldOperatorRowModel.FopName = req.FopName
	dbFieldOperatorRowModel.MobileNo = req.MobileNo
	dbFieldOperatorRowModel.EmailId = req.EmailId
	dbFieldOperatorRowModel.ShortDesc = req.ShortDesc
	dbFieldOperatorRowModel.FopState = req.FopState
	dbFieldOperatorRowModel.FopArea = req.FopArea

	dbErr, insertedId := dbaccess.Insert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbFieldOperatorRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new field operator.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	taskFieldOperatorAddedModel := &hktmodels.TaskFieldOperatorAddedRemovedOnSPModel{}
	taskFieldOperatorAddedModel.FopId = insertedId
	taskFieldOperatorAddedModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	isSendSuccess := repo.Instance().
		SendTaskToServer(hktconst.TASK_HKT_API_FIELD_OPERATOR_ADDED,
			service.ExeCtx.SessionToken, taskFieldOperatorAddedModel)

	if isSendSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New Field Operator Added succesfully")

	return true, addResponse
}

func (service FieldoperatorService) SelectById(fopID int64) (bool, interface{}) {

	dbErr, fopData := dbaccess.GetFieldOperatorById(service.ExeCtx.SessionInfo.Product.NodeDbConn, fopID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting field operator by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *fopData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched field operator info")
	return true, dbRecord[0]
}

func (service FieldoperatorService) GetFieldOperatorList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchFieldOperatorRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetFieldOperatorList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting field operator data list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched field operator list data.")

	return true, dataListResponse

}

func (service FieldoperatorService) Update(reqData *hktmodels.DBFieldOperatorUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, affectedRow := dbaccess.UpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating field operator info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Field operator updated successfully.")

	return true, nil
}

func (service FieldoperatorService) FieldOperatorShortDataList() (bool, interface{}) {

	cpmID := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, listData := dbaccess.GetFieldOperatorShortList(service.ExeCtx.SessionInfo.Product.NodeDbConn, cpmID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting fop short data list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Field Operator short data list.")

	return true, listData

}

func (service FieldoperatorService) FopSpAdd(req lmodels.APIFopSpAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBFopSpInsertRowModel{}
	dbRowModel.FopId = req.FopId
	dbRowModel.SpId = req.SpId
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, insertedId := dbaccess.FopSpInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while associating field operator with service point.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	taskFieldOperatorAddedRemovedOnSPModel := &hktmodels.TaskFieldOperatorAddedRemovedOnSPModel{}
	taskFieldOperatorAddedRemovedOnSPModel.FopId = req.FopId
	taskFieldOperatorAddedRemovedOnSPModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	isSendSuccess := repo.Instance().
		SendTaskToServer(hktconst.TASK_HKT_API_FIELD_OPERATOR_ADDED_ON_SP,
			service.ExeCtx.SessionToken, taskFieldOperatorAddedRemovedOnSPModel)

	if isSendSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Point associated with Field operater  succesfully")

	return true, addResponse
}

func (service FieldoperatorService) FopSpDelete(reqdata *lmodels.APIFopSpDeleteRequest) (isSuccess bool, successErrorData interface{}) {

	dbErr, affectedRow := dbaccess.FopSpDelete(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqdata)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while deassociating field operator with service point.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	taskFieldOperatorAddedRemovedOnSPModel := &hktmodels.TaskFieldOperatorAddedRemovedOnSPModel{}
	taskFieldOperatorAddedRemovedOnSPModel.FopId = reqdata.FopId
	taskFieldOperatorAddedRemovedOnSPModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	taskFieldOperatorAddedRemovedOnSPModel.SpId = reqdata.SpId

	isSendSuccess := repo.Instance().
		SendTaskToServer(hktconst.TASK_HKT_API_FIELD_OPERATOR_REMOVED_ON_SP,
			service.ExeCtx.SessionToken, taskFieldOperatorAddedRemovedOnSPModel)

	if isSendSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Point association with Field operater deleted successfully.")

	return true, nil
}

func (service FieldoperatorService) GetFopSpAssociation(fopID int64) (bool, interface{}) {

	dbErr, fopSpData := dbaccess.FopSpSelectByID(service.ExeCtx.SessionInfo.Product.NodeDbConn, fopID, service.ExeCtx.SessionInfo.Product.CustProdID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting fop-sp association.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *fopSpData

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched fopSp info")
	return true, dbRecord
}
