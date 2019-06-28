package models

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
)

func (r *APITaskHandlerModel) Convert(jsonData string) (bool, *gmodels.APITaskProcessModel) {

	apiTaskProcessModel := &gmodels.APITaskProcessModel{}
	apiTaskProcessModel.TaskPayload = r.PayloadType

	if isSuccess := ghelper.ConvertFromJSONString(jsonData, apiTaskProcessModel); isSuccess == false {
		logger.Context().WithField("JSONData", jsonData).Log("SPL.Server.Model.ConvertPayload", logger.Server, logger.Error, "Unable to convert json data")
		return false, nil
	}

	return true, apiTaskProcessModel
}
