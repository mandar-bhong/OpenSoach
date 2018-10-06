package job

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/vst/api/models"
	"opensoach.com/vst/api/webserver/job/dbaccess"
	constants "opensoach.com/vst/constants"
	hktmodels "opensoach.com/vst/models"
)

var SUB_MODULE_NAME = "VST.API.Job"

type JobService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service JobService) GetJobList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchJobRequestFilterDataModel)
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetJobList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting job list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched job list data.")

	return true, dataListResponse

}

func (service JobService) UpdateStatus(reqData *lmodels.APIJobStatusUpdateRequest) (isSuccess bool, successErrorData interface{}) {

	dbErr, jobData := dbaccess.GetJobDetailsByTokenId(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData.TokenId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting job details by token id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	jobDataRecords := *jobData
	jobDataItem := jobDataRecords[0]

	dbJobDeliveredTxnDataModel := hktmodels.DBJobDeliveredTxnDataModel{}
	dbJobDeliveredTxnDataModel.Tokenid = reqData.TokenId
	dbJobDeliveredTxnDataModel.BilledAmount = reqData.Amount

	isSuccess, jobDeliveredTxnDataJsonString := ghelper.ConvertToJSON(dbJobDeliveredTxnDataModel)
	if isSuccess == false {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Failed to convert to json")
		return false, nil
	}

	dbJobDeliveredTxnRowModel := hktmodels.DBJobDeliveredTxnRowModel{}
	dbJobDeliveredTxnRowModel.CpmId = jobDataItem.CpmId
	dbJobDeliveredTxnRowModel.FOPCode = jobDataItem.FopCode
	dbJobDeliveredTxnRowModel.ServiceInstanceID = jobDataItem.ServInId
	dbJobDeliveredTxnRowModel.TransactionDate = ghelper.GetCurrentTime()
	dbJobDeliveredTxnRowModel.Status = constants.JOB_TXN_STATUS_DELIVERED
	dbJobDeliveredTxnRowModel.TransactionData = jobDeliveredTxnDataJsonString

	dbTxErr, tx := dbaccess.GetDBTransaction(service.ExeCtx.SessionInfo.Product.NodeDbConn)

	if dbTxErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Transaction Error.", dbTxErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dberr, _ := dbaccess.InsertJobDeliveredTxn(tx, dbJobDeliveredTxnRowModel)
	if dberr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding job delivered txn record.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbJobStatusUpdateRowModel := &hktmodels.DBJobStatusUpdateRowModel{}
	dbJobStatusUpdateRowModel.TokenId = reqData.TokenId
	dbJobStatusUpdateRowModel.State = reqData.State

	dbErr, affectedRow := dbaccess.UpdateJobStatus(tx, dbJobStatusUpdateRowModel)
	if dbErr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating job status.", dbErr)

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

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Job status updated successfully.")

	return true, nil
}

func (service JobService) GetJobDetailsByTokenID(tokenID int64) (bool, interface{}) {

	dbErr, jobData := dbaccess.GetJobDetailsByTokenId(service.ExeCtx.SessionInfo.Product.NodeDbConn, tokenID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting job details by token id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *jobData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched job details")
	return true, jobData
}
