package processor

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	hktconst "opensoach.com/hpft/constants"
	hktmodels "opensoach.com/hpft/models"
	apitask "opensoach.com/hpft/server/processor/api"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

var apiTaskHandler map[string]pcmodels.APITaskProcessorHandlerModel

func init() {
	apiTaskHandler = make(map[string]pcmodels.APITaskProcessorHandlerModel)

	apiTaskHandler[hktconst.TASK_HKT_API_SERVICE_CONFIG_ADDED_ON_SP] = pcmodels.APITaskProcessorHandlerModel{Handler: apitask.ProcessSerConfigOnSP, PayloadType: &hktmodels.TaskSerConfigAddedOnSPModel{}}
	apiTaskHandler[hktconst.TASK_HKT_API_DEVICE_SP_ASSOCIATED] = pcmodels.APITaskProcessorHandlerModel{Handler: apitask.ProcessDeviceSPAssociated, PayloadType: &hktmodels.TaskSPDevAsscociatedModel{}}
	apiTaskHandler[hktconst.TASK_HKT_API_SERVICE_CONFIG_UPDATED] = pcmodels.APITaskProcessorHandlerModel{Handler: apitask.ProcessSerConfigUpdated, PayloadType: &hktmodels.TaskServConfigUpdatedModel{}}
	apiTaskHandler[hktconst.TASK_HKT_API_FIELD_OPERATOR_ADDED_ON_SP] = pcmodels.APITaskProcessorHandlerModel{Handler: apitask.ProcessFieldOperatorAddedOnSP, PayloadType: &hktmodels.TaskFieldOperatorAddedRemovedOnSPModel{}}
	apiTaskHandler[hktconst.TASK_HKT_API_FIELD_OPERATOR_REMOVED_ON_SP] = pcmodels.APITaskProcessorHandlerModel{Handler: apitask.ProcessFieldOperatorRemovedOnSP, PayloadType: &hktmodels.TaskFieldOperatorAddedRemovedOnSPModel{}}
	apiTaskHandler[hktconst.TASK_HKT_API_FIELD_OPERATOR_ADDED] = pcmodels.APITaskProcessorHandlerModel{Handler: apitask.ProcessFieldOperatorAdded, PayloadType: &hktmodels.TaskFieldOperatorAddedRemovedOnSPModel{}}
	apiTaskHandler[hktconst.TASK_HKT_API_PATIENT_STATUS_UPDATED] = pcmodels.APITaskProcessorHandlerModel{Handler: apitask.ProcessPatientStatusUpdated, PayloadType: &hktmodels.TaskPatientStatusUpdated{}}
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
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while executing api task", err)
		return "", err
	}

	if apiTaskResultModel.IsEPSync == true {
		ProcessEndPointTask(apiTaskResultModel.EPSyncData)
	}

	isJsonSuccess, jsonResultData := ghelper.ConvertToJSON(apiTaskResultModel)

	if isJsonSuccess == false {
		return "", fmt.Errorf("Unable to convert json data")
	}

	return jsonResultData, nil
}
