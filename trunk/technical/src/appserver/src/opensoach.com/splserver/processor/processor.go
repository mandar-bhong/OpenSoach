package processor

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/splserver/models"
)

var apiTaskHandler map[string]lmodels.APITaskHandlerModel

func init() {
	apiTaskHandler = make(map[string]lmodels.APITaskHandlerModel)
}

func RegisterHandler(handler map[string]interface{}) {
	handler[gmodels.TASK_HANDLER_API_SPL_CONTROLLER_KEY] = APITaskController
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

	err, apiTaskResultModel := taskHandlerModel.Handler(msg, apiTaskProcessModel.SessionToken,
		apiTaskProcessModel.TaskToken,
		apiTaskProcessModel.TaskPayload)

	if err != nil {
		return "", err
	}

	isJsonSuccess, jsonResultData := ghelper.ConvertToJSON(apiTaskResultModel)

	if isJsonSuccess == false {
		return "", fmt.Errorf("Unable to convert json data")
	}

	return jsonResultData, nil
}
