package corporate

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/corporate/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Corporate"

type CorporateService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (CorporateService) GetCorpDataList(corpListReqData lmodels.DataListRequest) (bool, interface{}) {

	dataListResponse := lmodels.DataListResponse{}

	filterModel := corpListReqData.Filter.(*lmodels.DBSearchCorpRequestFilterDataModel)

	CurrentPage := corpListReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * corpListReqData.Limit)

	dbErr, listData := dbaccess.GetCorpListData(repo.Instance().Context.Master.DBConn, filterModel, corpListReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched corporate list data.")

	return true, dataListResponse

}

func (CorporateService) GetCorpShortDataList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetCorpShortDataList(repo.Instance().Context.Master.DBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched corporate short data list.")

	return true, listData

}

func (service CorporateService) AddCorp(reqData *lmodels.DBSplCorpRowModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, insertedId := dbaccess.SplMasterCorpTableInsert(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	response := lmodels.RecordIdResponse{}
	response.RecId = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Corporate data added successfully.")

	return true, response
}

func (service CorporateService) UpdateCorp(reqData *lmodels.DBSplCorpRowModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, affectedRow := dbaccess.SplMasterCorpTableUpdate(repo.Instance().Context.Master.DBConn, reqData)
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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Corporate data updated successfully.")

	return true, nil
}
