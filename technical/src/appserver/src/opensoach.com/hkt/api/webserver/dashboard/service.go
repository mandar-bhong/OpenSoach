package dashboard

import (
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/dashboard/dbaccess"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
)

var SUB_MODULE_NAME = "HKT.API.Dashboard"

type DashboardService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service DashboardService) GetDeviceSummary() (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution device summary")

	dbErr, data := dbaccess.GetDeviceSummary(service.ExeCtx.SessionInfo.Product.NodeDbConn, service.ExeCtx.SessionInfo.Product.CustProdID)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured getting device summary.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	apiResponse := lmodels.APIDashboardDeviceSummaryResponse{}

	for _, dbDevSummaryDataModel := range data {

		apiResponse.TotalDevices = apiResponse.TotalDevices + dbDevSummaryDataModel.Count

		switch dbDevSummaryDataModel.ConnectionState {

		case pcconst.DB_DEVICE_CONNECTION_STATE_CONNECTED:
			apiResponse.Onlinedevices = dbDevSummaryDataModel.Count
		case pcconst.DB_DEVICE_CONNECTION_STATE_DISCONNECTED:
			apiResponse.Offlinedevices = dbDevSummaryDataModel.Count
		case pcconst.DB_DEVICE_CONNECTION_STATE_UNKNOWN:
		}
	}

	return true, apiResponse
}

func (service DashboardService) GetLocationSummary() (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution location summary")

	dbErr, data := dbaccess.GetLocationSummary(service.ExeCtx.SessionInfo.Product.NodeDbConn, service.ExeCtx.SessionInfo.Product.CustProdID)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured getting location summary.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	apiResponse := lmodels.APIDashboardLocationSummaryResponse{}

	for _, dbSummaryDataModel := range data {

		apiResponse.Total = apiResponse.Total + dbSummaryDataModel.Count

		switch dbSummaryDataModel.State {

		case pcconst.DB_SERVICE_POINT_STATE_ACTIVE:
			apiResponse.Active = dbSummaryDataModel.Count
		case pcconst.DB_SERVICE_POINT_STATE_INACTIVE:
		case pcconst.DB_SERVICE_POINT_STATE_SUSPENDED:
		}
	}

	return true, apiResponse

}
