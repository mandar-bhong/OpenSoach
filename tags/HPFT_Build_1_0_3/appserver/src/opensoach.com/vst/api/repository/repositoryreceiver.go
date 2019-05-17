package repository

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	hktconst "opensoach.com/vst/constants"
)

var SUB_MODULE_NAME = "SPL.Repo"

func (r *repo) SendTaskToServer(tasktag string, sessiontoken string, taskpayload interface{}) bool {

	apiTaskProcessModel := gmodels.APITaskProcessModel{}
	apiTaskProcessModel.TaskTag = tasktag
	apiTaskProcessModel.SessionToken = sessiontoken
	apiTaskProcessModel.TaskPayload = taskpayload

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(apiTaskProcessModel)

	if isJsonSuccess == false {
		logger.Context().WithField("API Task Model", apiTaskProcessModel).Log(SUB_MODULE_NAME, logger.Server, logger.Error, "Unable to convert json data for task processing")
		return false
	}

	err := r.ProdTaskContext.SubmitTask(hktconst.TASK_HANDLER_HKT_API_CONTROLLER, jsonData)

	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Unable to submit task for server processing", err)
		return false
	}

	return true
}
