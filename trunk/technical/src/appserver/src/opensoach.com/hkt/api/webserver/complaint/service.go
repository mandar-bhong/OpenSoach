package complaint

import (
	"time"

	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/complaint/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Complaint"

type ComplaintService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service ComplaintService) Add(req lmodels.APIComplaintAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBComplaintInsertRowModel{}
	dbRowModel.SpId = req.SpId
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.ComplaintTitle = req.ComplaintTitle
	dbRowModel.Description = req.Description
	dbRowModel.ComplaintBy = req.ComplaintBy
	dbRowModel.MobileNo = req.MobileNo
	dbRowModel.EmailId = req.EmailId
	dbRowModel.EmployeeId = req.EmployeeId
	dbRowModel.Severity = req.Severity
	dbRowModel.RaisedOn = time.Now()
	dbRowModel.ComplaintState = req.ComplaintState

	dbErr, insertedId := dbaccess.Insert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new complaint..", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New Complaint Added succesfully")

	return true, addResponse
}

func (service ComplaintService) Update(reqData *hktmodels.DBComplaintUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, affectedRow := dbaccess.UpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating complaint", dbErr)

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Complaint updated successfully.")

	return true, nil
}

func (service ComplaintService) SelectById(complaintID int64) (bool, interface{}) {

	dbErr, complaintData := dbaccess.ComplaintTableSelectByID(service.ExeCtx.SessionInfo.Product.NodeDbConn, complaintID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while fetching complaint info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *complaintData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Complaint info")
	return true, dbRecord[0]
}

func (service ComplaintService) ComplaintList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchComplaintRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetComplaintList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while fetching complaint list data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched complaint list data.")

	return true, dataListResponse

}
