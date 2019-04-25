package dashboard

import (
	"github.com/gin-gonic/gin"

	"opensoach.com/core/logger"
	"opensoach.com/hpft/api/constants"
	lhelper "opensoach.com/hpft/api/helper"
	lmodels "opensoach.com/hpft/api/models"
	repo "opensoach.com/hpft/api/repository"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_DASHBOARD_DEVICE_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_LOCATION_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_FEEDBACK_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_TASK_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_COMPLAINT_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_FEEDBACKS_PER_MONTH, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_COMPLAINTS_PER_MONTH, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_TOP_COMPLAINTS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_TASK_PER_MONTH, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_TOP_FEEDBACKS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_PATIENT_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_PATIENT_HOSPITALIZED_PER_MONTH, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_DASHBOARD_DEVICE_SUMMARY:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetDeviceSummary()

		break

	case constants.API_DASHBOARD_LOCATION_SUMMARY:
		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetLocationSummary()
		break

	case constants.API_DASHBOARD_FEEDBACK_SUMMARY:

		feedbackReq := lmodels.APIDashboardFeedbackFilterModel{}
		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &feedbackReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetFeedbackSummary(feedbackReq)
		break

	case constants.API_DASHBOARD_TASK_SUMMARY:

		taskReq := lmodels.APIDashboardTaskRequest{}
		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &taskReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetTaskSummary(taskReq)
		break

	case constants.API_DASHBOARD_COMPLAINT_SUMMARY:
		complaintReq := lmodels.APIDashboardComplaintFilterModel{}
		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &complaintReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetComplaintSummary(complaintReq)

		break

	case constants.API_DASHBOARD_FEEDBACKS_PER_MONTH:

		req := lmodels.APIFeedbacksPerMonthRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.FeedbackPerMonth(req)

		break

	case constants.API_DASHBOARD_COMPLAINTS_PER_MONTH:

		req := lmodels.APIComplaintsByMonthRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.NoOfComplaints(req)

		break

	case constants.API_DASHBOARD_TOP_COMPLAINTS:

		req := lmodels.APITopActiveComplaintsRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.TopComplaints(req)

		break

	case constants.API_DASHBOARD_TASK_PER_MONTH:

		req := lmodels.APITaskByMonthRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.TaskSummaryPerMonth(req)

		break

	case constants.API_DASHBOARD_TOP_FEEDBACKS:

		req := lmodels.APITopFeedbacksRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.TopFeedbacks(req)

		break

	case constants.API_DASHBOARD_PATIENT_SUMMARY:
		patientReq := lmodels.APIDashboardPatientFilterModel{}
		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &patientReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetPatientSummary(patientReq)

		break

	case constants.API_DASHBOARD_PATIENT_HOSPITALIZED_PER_MONTH:

		req := lmodels.APIPatientHospitalisedByMonthRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.PatientHospitalisedSummaryPerMonth(req)

		break

	}

	return isSuccess, resultData
}
