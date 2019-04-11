package vehicle

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/vst/api/webserver/vehicle/dbaccess"
)

var SUB_MODULE_NAME = "VST.API.Vehicle"

type VehicleService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service VehicleService) SelectById(vehicleID int64) (bool, interface{}) {

	dbErr, vehicleData := dbaccess.VehicleTableSelectByID(service.ExeCtx.SessionInfo.Product.NodeDbConn, vehicleID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while fetching vehicle info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *vehicleData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched vehicle info")
	return true, dbRecord[0]
}
