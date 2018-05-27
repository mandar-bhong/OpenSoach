package manager

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/prodcore/dbaccess"
)

var SUB_MODULE_NAME = "ProdCore.Manager"

func GetTaskExecutionContext(mstdbconn string, cpmid int64) (bool, *gmodels.TaskExeContextModel) {

	dbErr, data := dbaccess.GetCPMDBDetails(mstdbconn, cpmid)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while get db details by cpmid", dbErr)
		return false, data
	}

	return true, data

}
