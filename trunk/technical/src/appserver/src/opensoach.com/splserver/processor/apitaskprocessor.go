package processor

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	dbaccess "opensoach.com/splserver/dbaccess"
	lmodels "opensoach.com/splserver/models"
	repo "opensoach.com/splserver/repository"
)

const SUB_MODULE_NAME = "Splserver.Processor"

func APIHandlerCustProdAssociated(msg string, sessionkey string,
	tasktoken string,
	taskData interface{}) (error, lmodels.APITaskResultModel) {

	result := lmodels.APITaskResultModel{}

	taskAPICustProdAssociatedModel := taskData.(*gmodels.TaskAPICustProdAssociatedModel)

	err, dbConn := dbaccess.GetDBConnectionByID(repo.Instance().Context.Master.DBConn, taskAPICustProdAssociatedModel.DbiId)

	if err != nil {
		//Error need to retry
		logger.Context().WithField("Task Data", taskData).
			WithField("TaskToken", tasktoken).
			WithField("DBConn", repo.Instance().Context.Master.DBConn).
			WithField("TaskExecData", taskAPICustProdAssociatedModel).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to get instance dbconn.", err)

		errModel := lmodels.APITaskResultErrorDataModel{}
		errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

		result.IsSuccess = false
		result.ErrorData = errModel

		return err, result
	}

	dbInstanceCpmIdInsert := &lmodels.APIDBInstanceCpmIdInsertModel{}
	dbInstanceCpmIdInsert.CpmId = taskAPICustProdAssociatedModel.CpmId

	dbErr, insertedId := dbaccess.UpdateCPMIDToInstDB(dbConn, dbInstanceCpmIdInsert)

	if dbErr != nil {
		logger.Context().WithField("Task Data", taskData).
			WithField("TaskToken", tasktoken).
			WithField("DBConn", dbConn).
			WithField("TaskExecData", dbInstanceCpmIdInsert).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to update cpm id in instance db.", err)

		errModel := lmodels.APITaskResultErrorDataModel{}
		errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

		result.IsSuccess = false
		result.ErrorData = errModel

		return err, result

	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully inserted cpm id in instance db.")

	apiRecordResponse := gmodels.APIRecordAddResponse{}
	apiRecordResponse.RecordID = insertedId

	result.IsSuccess = true
	result.Data = apiRecordResponse

	return nil, result

}

//taskAPICustProdAssociatedModel.DbiId = reqData.DbiId

//	if isSuccess := repo.Instance().TaskQue.
//		SubmitTask(gmodels.TASK_API_CUST_PROD_ASSOCIATED,
//			taskAPICustProdAssociatedModel); isSuccess == false {
//		logger.Context().Log(SUB_MODULE_NAME, logger.Error, logger.Normal, "Error occured while submiting task for cust prod assoc")
//	}
