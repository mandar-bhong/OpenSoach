package processor

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	hktmodels "opensoach.com/hkt/models"
	apitask "opensoach.com/hkt/server/processor/api"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

var apiTaskHandler map[string]pcmodels.APITaskHandlerModel

func init() {
	apiTaskHandler = make(map[string]pcmodels.APITaskHandlerModel)

	apiTaskHandler[gmodels.TASK_API_CUST_PROD_ASSOCIATED] = pcmodels.APITaskHandlerModel{Handler: apitask.ProcessSerConfigOnSP, PayloadType: &hktmodels.TaskSerConfigAddedOnSPModel{}}
}

func APITaskController(msg string) (string, error) {

	apiTaskProcessModel := &gmodels.APITaskProcessModel{}

	if isSuccess := ghelper.ConvertFromJSONString(msg, apiTaskProcessModel); isSuccess == false {

		//Error condition
		return "", nil
	}

	taskHandlerModel, hasItem := apiTaskHandler[apiTaskProcessModel.TaskTag]

	if hasItem == false {
		//Handler is not attached
		return "", nil
	}

	isSuccess, apiTaskProcessModel := taskHandlerModel.Convert(msg)

	if isSuccess == false {
		return "", nil
	}

	apiTaskExecutionCtx := &pcmodels.APITaskExecutionCtx{}
	apiTaskExecutionCtx.Message = msg
	apiTaskExecutionCtx.SessionKey = apiTaskProcessModel.SessionToken
	apiTaskExecutionCtx.Token = apiTaskProcessModel.TaskToken
	apiTaskExecutionCtx.TaskData = apiTaskProcessModel.TaskPayload

	err, apiTaskResultModel := taskHandlerModel.Handler(apiTaskExecutionCtx)

	if err != nil {
		return "", err
	}

	isJsonSuccess, jsonResultData := ghelper.ConvertToJSON(apiTaskResultModel)

	if isJsonSuccess == false {
		return "", fmt.Errorf("Unable to convert json data")
	}

	return jsonResultData, nil
}
