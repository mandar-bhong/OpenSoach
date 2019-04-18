package services

import (
	"errors"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	pcconst "opensoach.com/prodcore/constants"
)

var SUB_MODULE_NAME = "ProdCore.Services"

//DtcollOnlineDevicesService This service get online devices token
type DtcollOnlineDevicesService struct {
	*ServiceContext
	NextHandler IHandler
}

//Handle Get online device token from central cache
func (r *DtcollOnlineDevicesService) Handle(serctx *ServiceContext) error {

	getErr, tokenlistjsonstring := r.Repo.ProdTaskContext.ProcessTask(pcconst.TASK_GET_ONLINE_DEVICES, "")
	if getErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while submitting task get online devices.", getErr)
		return getErr
	}

	isJSONSuccess := ghelper.ConvertFromJSONString(tokenlistjsonstring, &r.ServiceRuntime.Tokens)

	if isJSONSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert json device packet ", nil)
		return errors.New("failed to convert json")
	}

	logger.Context().WithField("online devices", tokenlistjsonstring).LogDebug(SUB_MODULE_NAME, logger.Normal, "Online devices")

	if r.NextHandler != nil {
		err := r.NextHandler.Handle(serctx)
		return err
	}

	return nil

}
