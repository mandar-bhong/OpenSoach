package processor

import (
	"strconv"

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

	dbErr, spcmasterdata := dbaccess.GetDBHktMasterSpCategory(repo.Instance().Context.ProdMst.DBConn)

	if dbErr != nil {
		logger.Context().WithField("Task Data", taskData).
			WithField("TaskToken", tasktoken).
			WithField("DBConn", repo.Instance().Context.ProdMst.DBConn).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to get instance dbconn.", dbErr)

		errModel := lmodels.APITaskResultErrorDataModel{}
		errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

		result.IsSuccess = false
		result.ErrorData = errModel

		return dbErr, result
	}

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

	dbInstanceCpmIdInsert := &lmodels.APITaskDBInstanceCpmIdInsertModel{}
	dbInstanceCpmIdInsert.CpmId = taskAPICustProdAssociatedModel.CpmId

	var isDBOpSuccess = true

	dbTxErr, tx := dbaccess.GetDBTransaction(dbConn)

	if dbTxErr != nil {
		errModel := lmodels.APITaskResultErrorDataModel{}
		errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE
		result.IsSuccess = false
		result.ErrorData = errModel

		return err, result
	}

	dberr, _ := dbaccess.UpdateCPMIDToInstDB(tx, dbInstanceCpmIdInsert)

	if dberr != nil {

		isDBOpSuccess = false

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

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

	for _, spcmasteritem := range *spcmasterdata {

		dbInstanceSpCategoryInsertModel := &lmodels.APITaskDBInstanceSpCategoryInsertModel{}
		dbInstanceSpCategoryInsertModel.CpmId = taskAPICustProdAssociatedModel.CpmId
		dbInstanceSpCategoryInsertModel.SpcId = spcmasteritem.SpcId
		dbInstanceSpCategoryInsertModel.SpcName = spcmasteritem.SpcName
		dbInstanceSpCategoryInsertModel.ShortDesc = spcmasteritem.ShortDesc

		dbErr, _ := dbaccess.UpdateSpCategoryToInstanceDB(tx, dbInstanceSpCategoryInsertModel)

		if dbErr != nil {

			isDBOpSuccess = false

			txErr := tx.Rollback()

			if txErr != nil {
				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
			}

			logger.Context().WithField("Task Data", taskData).
				WithField("TaskToken", tasktoken).
				WithField("DBConn", repo.Instance().Context.Master.DBConn).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to get instance dbconn.", dbErr)
			isDBOpSuccess = false

			break
		}
	}

	if isDBOpSuccess {

		txErr := tx.Commit()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
			errModel := lmodels.APITaskResultErrorDataModel{}
			errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

			result.IsSuccess = false
			result.ErrorData = errModel
			return nil, result
		}

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully Inserted cpm id and sp category in instance db")

		result.IsSuccess = true

		return nil, result

	} else {
		txErr := tx.Rollback()
		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}
		errModel := lmodels.APITaskResultErrorDataModel{}
		errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

		result.IsSuccess = false
		result.ErrorData = errModel

		return nil, result
	}
}

func APIHandlerDevProdAssociated(msg string, sessionkey string,
	tasktoken string,
	taskData interface{}) (error, lmodels.APITaskResultModel) {

	result := lmodels.APITaskResultModel{}

	taskDevProdAsscociatedModel := taskData.(*gmodels.TaskDevProdAsscociatedModel)

	err, dbConn := dbaccess.GetDBConnectionByCpmID(repo.Instance().Context.Master.DBConn, taskDevProdAsscociatedModel.CpmId)

	if err != nil {
		//Error need to retry
		logger.Context().WithField("Task Data", taskData).
			WithField("TaskToken", tasktoken).
			WithField("DBConn", repo.Instance().Context.Master.DBConn).
			WithField("TaskExecData", taskDevProdAsscociatedModel).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to get instance dbconn.", err)

		errModel := lmodels.APITaskResultErrorDataModel{}
		errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

		result.IsSuccess = false
		result.ErrorData = errModel

		return err, result
	}

	dbInstanceDevInsertModel := &lmodels.APITaskDBInstanceDevInsertRowModel{}
	dbInstanceDevInsertModel.CpmId = taskDevProdAsscociatedModel.CpmId
	dbInstanceDevInsertModel.DevId = taskDevProdAsscociatedModel.DevId
	dbInstanceDevInsertModel.Serialno = taskDevProdAsscociatedModel.Serialno
	dbInstanceDevInsertModel.DevName = "device " + strconv.FormatInt(taskDevProdAsscociatedModel.DevId, 10)

	dbErr, insertedId := dbaccess.UpdateDevToInstDB(dbConn, dbInstanceDevInsertModel)

	if dbErr != nil {
		logger.Context().WithField("Task Data", taskData).
			WithField("TaskToken", tasktoken).
			WithField("DBConn", dbConn).
			WithField("TaskExecData", dbInstanceDevInsertModel).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to update dev data in instance db.", err)

		errModel := lmodels.APITaskResultErrorDataModel{}
		errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

		result.IsSuccess = false
		result.ErrorData = errModel

		return err, result

	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully inserted dev data in instance db.")

	apiRecordResponse := gmodels.APIRecordAddResponse{}
	apiRecordResponse.RecordID = insertedId

	result.IsSuccess = true
	result.Data = apiRecordResponse

	return nil, result

}

func APIHandlerCustServPointAssociated(msg string, sessionkey string,
	tasktoken string,
	taskData interface{}) (error, lmodels.APITaskResultModel) {

	result := lmodels.APITaskResultModel{}

	taskCustServicePointAssociatedModel := taskData.(*gmodels.TaskCustServicePointAssociatedModel)

	err, dbConn := dbaccess.GetDBConnectionByCpmID(repo.Instance().Context.Master.DBConn, taskCustServicePointAssociatedModel.CpmId)

	if err != nil {
		//Error need to retry
		logger.Context().WithField("Task Data", taskData).
			WithField("TaskToken", tasktoken).
			WithField("DBConn", repo.Instance().Context.Master.DBConn).
			WithField("TaskExecData", taskCustServicePointAssociatedModel).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to get instance dbconn.", err)

		errModel := lmodels.APITaskResultErrorDataModel{}
		errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

		result.IsSuccess = false
		result.ErrorData = errModel

		return err, result
	}

	for i := 0; i < len(taskCustServicePointAssociatedModel.SpIdList); i++ {

		dbInstanceSpInsertModel := &lmodels.APITaskDBNodeSpInsertRowModel{}
		dbInstanceSpInsertModel.SpId = taskCustServicePointAssociatedModel.SpIdList[i]
		dbInstanceSpInsertModel.SpcId = 0
		dbInstanceSpInsertModel.CpmId = taskCustServicePointAssociatedModel.CpmId
		Spid := dbInstanceSpInsertModel.SpId
		dbInstanceSpInsertModel.SpName = "Service Point " + strconv.FormatInt(Spid, 10)
		dbInstanceSpInsertModel.SpState = taskCustServicePointAssociatedModel.SpState
		dbInstanceSpInsertModel.SpStateSince = taskCustServicePointAssociatedModel.SpStateSince

		dbErr, _ := dbaccess.UpdateServicePointsToInstDB(dbConn, dbInstanceSpInsertModel)

		if dbErr != nil {
			logger.Context().WithField("Task Data", taskData).
				WithField("TaskToken", tasktoken).
				WithField("DBConn", dbConn).
				WithField("TaskExecData", dbInstanceSpInsertModel).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to update service points to instance db.", err)

			//TODO: Need to handle only failure ids
			// errModel := lmodels.APITaskResultErrorDataModel{}
			// errModel.ErrorCode = gmodels.MOD_TASK_OPER_ERR_DATABASE

			// result.IsSuccess = false
			// result.ErrorData = errModel

			//return err, result
		}

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully inserted dev data in instance db.")

		result.IsSuccess = true
	}
	return nil, result
}
