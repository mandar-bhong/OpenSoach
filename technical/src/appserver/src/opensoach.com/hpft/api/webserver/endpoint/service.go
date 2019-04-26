package endpoint

import (
	"opensoach.com/core/logger"
	"opensoach.com/hpft/api/webserver/endpoint/dbaccess"
	hpftmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.Endpoint"

type EndpointService struct {
	ExeCtx *gmodels.DeviceUserExecutionContext
}

func (service EndpointService) GetPatientAdmissionList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	usrID := service.ExeCtx.DeviceUserSessionInfo.UserID
	filterModel := listReqData.Filter.(*hpftmodels.DBDeviceSearchPatientRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.DeviceUserSessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientList(service.ExeCtx.DeviceUserSessionInfo.Product.NodeDbConn, usrID, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient admission list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient admission list data.")

	return true, dataListResponse

}
