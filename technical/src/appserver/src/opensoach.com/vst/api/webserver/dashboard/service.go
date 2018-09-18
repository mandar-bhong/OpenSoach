package dashboard

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	lmodels "opensoach.com/vst/api/models"
	"opensoach.com/vst/api/webserver/dashboard/dbaccess"
	"opensoach.com/vst/constants"
	hktconst "opensoach.com/vst/constants"
	hktmodels "opensoach.com/vst/models"
)

var SUB_MODULE_NAME = "VST.API.Dashboard"

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
		}
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched device summary")

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

	dberr, inUseCount := dbaccess.GetInUseLocations(service.ExeCtx.SessionInfo.Product.NodeDbConn, service.ExeCtx.SessionInfo.Product.CustProdID)

	if dberr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured getting location summary.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	apiResponse := lmodels.APIDashboardLocationSummaryResponse{}

	apiResponse.InUse = inUseCount[0].Count

	for _, dbSummaryDataModel := range data {

		apiResponse.Total = apiResponse.Total + dbSummaryDataModel.Count

		switch dbSummaryDataModel.State {

		case pcconst.DB_SERVICE_POINT_STATE_ACTIVE:
			apiResponse.Active = dbSummaryDataModel.Count
		case pcconst.DB_SERVICE_POINT_STATE_INACTIVE:
		case pcconst.DB_SERVICE_POINT_STATE_SUSPENDED:
		}
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched location summary")

	return true, apiResponse

}

func (service DashboardService) GetFeedbackSummary(req lmodels.APIDashboardFeedbackFilterModel) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution feedback summary")

	req.CPMID = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, data := dbaccess.GetFeedbackSummary(service.ExeCtx.SessionInfo.Product.NodeDbConn, req)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured getting feedback summary.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	apiResponse := lmodels.APIDashboardFeedbackResponse{}

	for _, dbSummaryDataModel := range data {

		switch dbSummaryDataModel.Feedback {
		case hktconst.DB_FEEDBACK_RATING_1:
			apiResponse.Rating1 = dbSummaryDataModel.Count
		case hktconst.DB_FEEDBACK_RATING_2:
			apiResponse.Rating2 = dbSummaryDataModel.Count
		case hktconst.DB_FEEDBACK_RATING_3:
			apiResponse.Rating3 = dbSummaryDataModel.Count
		case hktconst.DB_FEEDBACK_RATING_4:
			apiResponse.Rating4 = dbSummaryDataModel.Count
		case hktconst.DB_FEEDBACK_RATING_5:
			apiResponse.Rating5 = dbSummaryDataModel.Count
		}
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched feedback summary")

	return true, apiResponse

}

func (service DashboardService) GetTaskSummary(req lmodels.APIDashboardTaskRequest) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution task summary")

	filterModel := hktmodels.DBTaskSummaryFilterDataModel{}
	filterModel.SpId = req.SPId
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, data := dbaccess.GetTaskSummary(service.ExeCtx.SessionInfo.Product.NodeDbConn, req, filterModel)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured getting task summary.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	apiResponse := lmodels.APIDashboardTaskResponse{}

	for _, dbSummaryDataModel := range data {

		switch dbSummaryDataModel.Status {
		case hktconst.DB_TASK_ONTIME:
			apiResponse.Ontime = dbSummaryDataModel.Count
		case hktconst.DB_TASK_DELAYED:
			apiResponse.Delayed = dbSummaryDataModel.Count
		}
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched task summary")

	return true, apiResponse

}

func (service DashboardService) GetComplaintSummary(req lmodels.APIDashboardComplaintFilterModel) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution complaint summary")

	req.CPMID = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, data := dbaccess.GetComplaintSummary(service.ExeCtx.SessionInfo.Product.NodeDbConn, req)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured getting complaint summary.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	apiResponse := lmodels.APIDashboardComplaintResponse{}

	for _, dbSummaryDataModel := range data {

		switch dbSummaryDataModel.ComplaintState {
		case hktconst.DB_COMPLAINT_STATE_OPEN:
			apiResponse.Open = dbSummaryDataModel.Count
		case hktconst.DB_COMPLAINT_STATE_CLOSED:
			apiResponse.Close = dbSummaryDataModel.Count
		case hktconst.DB_COMPLAINT_STATE_INPROGRESS:
			apiResponse.Inprogress = dbSummaryDataModel.Count

		}
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched complaint summary")

	return true, apiResponse

}

func (service DashboardService) FeedbackPerMonth(req lmodels.APIFeedbacksPerMonthRequest) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution feedback per month")

	filterModel := hktmodels.DBFeedbacksPerMonthFilterDataModel{}
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	filterModel.SpId = req.SpID

	dbErr, feedbackList := dbaccess.GetFeedbackPerMonth(service.ExeCtx.SessionInfo.Product.NodeDbConn, req, filterModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting feedback per month.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched feedback per month")
	return true, feedbackList
}

func (service DashboardService) NoOfComplaints(req lmodels.APIComplaintsByMonthRequest) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution complaints per month")

	filterModel := hktmodels.DBNoOfComplaintsPerMonthsFilterDataModel{}
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	filterModel.SpId = req.SpID

	dbErr, complaintList := dbaccess.GetNoOfComplaintsPerMonth(service.ExeCtx.SessionInfo.Product.NodeDbConn, req, filterModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting complaints per month.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched complaints per month")
	return true, complaintList
}

func (service DashboardService) TopComplaints(req lmodels.APITopActiveComplaintsRequest) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution topo complaints.")

	filterModel := hktmodels.DBTopComplaintsFilterDataModel{}
	filterModel.ComplaintState = constants.DB_COMPLAINT_STATE_OPEN
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	filterModel.SpId = req.SpID

	dbErr, complaintList := dbaccess.SelectTopComplaints(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, req.NoOfComplaints)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting top complaints.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched top active complaints")
	return true, complaintList
}

func (service DashboardService) TaskSummaryPerMonth(req lmodels.APITaskByMonthRequest) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution task summary per month")

	filterModel := hktmodels.DBTaskPerMonthFilterDataModel{}
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	filterModel.SpId = req.SpID

	dbErr, taskList := dbaccess.GetTaskSummaryPerMonth(service.ExeCtx.SessionInfo.Product.NodeDbConn, req, filterModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while gettting task summary per month.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched task summary per month.")
	return true, taskList
}

func (service DashboardService) TopFeedbacks(req lmodels.APITopFeedbacksRequest) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution top feedbacks.")

	filterModel := hktmodels.DBTopFeedbackFilterDataModel{}
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	filterModel.SpId = req.SpID

	dbErr, dataList := dbaccess.SelectTopFeedbacks(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, req.NoOfFeedbacks)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting top feedbacks.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched top feedbacks")
	return true, dataList
}

func (service DashboardService) SnapshotData(req lmodels.APIDashboardVehicleRequest) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution snapshot data.")

	cpmId := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, dataList := dbaccess.GetSnapshotData(service.ExeCtx.SessionInfo.Product.NodeDbConn, req, cpmId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting snapshot data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched snapshot data")
	return true, dataList
}

func (service DashboardService) AverageTime(req lmodels.APIDashboardVehicleRequest) (bool, interface{}) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution snapshot data.")

	dbErr, dataList := dbaccess.GetVhlAverageTime(service.ExeCtx.SessionInfo.Product.NodeDbConn, req)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting average time data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched snapshot data")
	return true, dataList[0]
}
