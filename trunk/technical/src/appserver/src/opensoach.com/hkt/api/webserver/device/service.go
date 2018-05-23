package device

import (
	"opensoach.com/core/logger"
	"opensoach.com/hkt/api/webserver/device/dbaccess"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Device"

type DeviceService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service DeviceService) DeviceShortList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetDeviceShortDataList(service.ExeCtx.SessionInfo.Product.NodeDbConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched device short data list.")

	return true, listData

}
