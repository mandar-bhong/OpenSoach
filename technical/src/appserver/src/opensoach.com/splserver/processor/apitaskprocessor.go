package processor

import (
	gmodels "opensoach.com/models"
	"opensoach.com/splserver/dbaccess"
)

func APIHandlerCustProdAssociated(msg string, sessionkey string,
	tasktoken string,
	taskData interface{}) {

	taskAPICustProdAssociatedModel := taskData.(*gmodels.TaskAPICustProdAssociatedModel)

	err, dbConn := dbaccess.GetDBConnectionByID(taskAPICustProdAssociatedModel.DbiId)

	if err != nil {
		//Error need to retry
	}

}

//taskAPICustProdAssociatedModel.DbiId = reqData.DbiId

//	if isSuccess := repo.Instance().TaskQue.
//		SubmitTask(gmodels.TASK_API_CUST_PROD_ASSOCIATED,
//			taskAPICustProdAssociatedModel); isSuccess == false {
//		logger.Context().Log(SUB_MODULE_NAME, logger.Error, logger.Normal, "Error occured while submiting task for cust prod assoc")
//	}
