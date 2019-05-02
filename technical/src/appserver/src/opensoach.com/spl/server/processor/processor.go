package processor

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/server/models"
)

var apiTaskHandler map[string]lmodels.APITaskHandlerModel

func init() {
	apiTaskHandler = make(map[string]lmodels.APITaskHandlerModel)
	apiTaskHandler[gmodels.TASK_API_CUST_PROD_ASSOCIATED] = lmodels.APITaskHandlerModel{Handler: APIHandlerCustProdAssociated, PayloadType: &gmodels.TaskAPICustProdAssociatedModel{}}
	apiTaskHandler[gmodels.TASK_API_DEV_PROD_ASSOCIATED] = lmodels.APITaskHandlerModel{Handler: APIHandlerDevProdAssociated, PayloadType: &gmodels.TaskDevProdAsscociatedModel{}}
	apiTaskHandler[gmodels.TASK_API_CUST_SERVICE_POINT_ASSOCIATED] = lmodels.APITaskHandlerModel{Handler: APIHandlerCustServPointAssociated, PayloadType: &gmodels.TaskCustServicePointAssociatedModel{}}
	apiTaskHandler[gmodels.TASK_API_USER_ASSOCIATED] = lmodels.APITaskHandlerModel{Handler: APIHandlerUserAssociated, PayloadType: &gmodels.TaskUserAssociatedModel{}}
	apiTaskHandler[gmodels.TASK_API_USER_SEND_OTP_EMAIL_NOTIFICATION] = lmodels.APITaskHandlerModel{Handler: APIHandlerSendOTPEmailNotification, PayloadType: &gmodels.TaskUserForgotPasswordInfoModel{}}
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
