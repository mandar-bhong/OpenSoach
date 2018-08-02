package dbaccess

import (
	"fmt"
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hkthelper "opensoach.com/hpft/api/helper"
	lmodels "opensoach.com/hpft/api/models"
	"opensoach.com/hpft/constants/dbquery"
	hktmodels "opensoach.com/hpft/models"
	pcconst "opensoach.com/prodcore/constants"
)

var SUB_MODULE_NAME = "HPFT.API.Dashboard.DB"

func GetDeviceSummary(dbConn string, cpmid int64) (error, []hktmodels.DBDashBoardDeviceSummaryDataModel) {
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceSummary")
	data := []hktmodels.DBDashBoardDeviceSummaryDataModel{}

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = dbquery.QUERY_SPL_NODE_DASHBOARD_DEVICE_SUMMARY
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select(cpmid)
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func GetLocationSummary(dbConn string, cpmid int64) (error, []hktmodels.DBDashBoardLocationSummaryDataModel) {
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetLocationSummary")
	data := []hktmodels.DBDashBoardLocationSummaryDataModel{}

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = dbquery.QUERY_SPL_NODE_DASHBOARD_LOCATION_SUMMARY
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select(cpmid)
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func GetFeedbackSummary(dbConn string, req lmodels.APIDashboardFeedbackFilterModel) (error, []hktmodels.DBDashBoardFeedbackDataModel) {
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetFeedbackSummary")
	data := []hktmodels.DBDashBoardFeedbackDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(req)

	if req.StartTime != nil && req.EndTime != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := req.StartTime.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := req.EndTime.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " raised_on between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_SPL_NODE_DASHBOARD_FEEDBACK, "$WhereCondition$", whereCondition, 1)

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func GetTaskSummary(dbConn string, req lmodels.APIDashboardTaskRequest, filtermodel hktmodels.DBTaskSummaryFilterDataModel) (error, []hktmodels.DBDashBoardTaskDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetTaskSummary")

	data := []hktmodels.DBDashBoardTaskDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if req.StartTime != nil && req.EndTime != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := req.StartTime.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := req.EndTime.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " txn_date between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_SPL_NODE_DASHBOARD_TASK, "$WhereCondition$", whereCondition, 1)

	fmt.Println(query)

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func GetComplaintSummary(dbConn string, req lmodels.APIDashboardComplaintFilterModel) (error, []hktmodels.DBDashBoardComplaintDataModel) {
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetFeedbackSummary")
	data := []hktmodels.DBDashBoardComplaintDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(req)

	if req.StartTime != nil && req.EndTime != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := req.StartTime.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := req.EndTime.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " raised_on between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_SPL_NODE_DASHBOARD_COMPLAINT_SUMMARY, "$WhereCondition$", whereCondition, 1)

	logger.Context().WithField("Query", query).LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution query")

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func GetInUseLocations(dbConn string, cpmid int64) (error, []hktmodels.DBDashBoardInUseLocationDataModel) {
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetLocationSummary")
	data := []hktmodels.DBDashBoardInUseLocationDataModel{}

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = dbquery.QUERY_SPL_NODE_DASHBOARD_IN_USE_LOCATION_COUNT
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select(cpmid)
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func GetFeedbackPerMonth(dbConn string, req lmodels.APIFeedbacksPerMonthRequest, filtermodel hktmodels.DBFeedbacksPerMonthFilterDataModel) (error, []hktmodels.DBFeedbacksPerMonthDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetFeedbackPerMonth")

	data := []hktmodels.DBFeedbacksPerMonthDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if req.StartDate != nil && req.EndDate != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := req.StartDate.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := req.EndDate.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " raised_on between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_FEEDBACKS_PER_MONTH, "$WhereCondition$", whereCondition, 1)

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func GetNoOfComplaintsPerMonth(dbConn string, req lmodels.APIComplaintsByMonthRequest, filtermodel hktmodels.DBNoOfComplaintsPerMonthsFilterDataModel) (error, []hktmodels.DBNoOfComplaintsPerMonthDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetNoOfComplaintsPerMonth")

	data := []hktmodels.DBNoOfComplaintsPerMonthDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if req.StartDate != nil && req.EndDate != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := req.StartDate.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := req.EndDate.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " raised_on between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_NO_OF_COMPLAINTS_PER_MONTH, "$WhereCondition$", whereCondition, 1)

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func SelectTopComplaints(dbConn string, filtermodel hktmodels.DBTopComplaintsFilterDataModel, noofcomplaints int) (error, *[]lmodels.APITopActiveComplaintsResponse) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SelectTopComplaints")

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	Query := strings.Replace(dbquery.QUERY_GET_TOP_COMPLAINTS, "$WhereCondition$", whereCondition, 1)

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.APITopActiveComplaintsResponse{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.Query = Query
	selErr := selDBCtx.Select(noofcomplaints)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetTaskSummaryPerMonth(dbConn string, req lmodels.APITaskByMonthRequest, filtermodel hktmodels.DBTaskPerMonthFilterDataModel) (error, []hktmodels.DBTaskSummaryPerMonthDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetTaskSummaryPerMonth")

	data := []hktmodels.DBTaskSummaryPerMonthDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if req.StartDate != nil && req.EndDate != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := req.StartDate.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := req.EndDate.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " txn_date between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_TASK_SUMMARY_PER_MONTH, "$WhereCondition$", whereCondition, 1)

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}

func SelectTopFeedbacks(dbConn string, filtermodel hktmodels.DBTopFeedbackFilterDataModel, nooffeedbacks int) (error, *[]lmodels.APITopFeedbacksResponse) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SelectTopFeedbacks")

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition + " and feedback_comment is not null"
	}

	Query := strings.Replace(dbquery.QUERY_GET_TOP_FEEDBACKS, "$WhereCondition$", whereCondition, 1)

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.APITopFeedbacksResponse{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.Query = Query
	selErr := selDBCtx.Select(nooffeedbacks)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetPatientSummary(dbConn string, req lmodels.APIDashboardPatientFilterModel) (error, []hktmodels.DBDashBoardPatientDataModel) {
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientSummary")
	data := []hktmodels.DBDashBoardPatientDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(req)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_SPL_NODE_DASHBOARD_PATIENT_SUMMARY, "$WhereCondition$", whereCondition, 1)

	logger.Context().WithField("Query", query).LogDebug(SUB_MODULE_NAME, logger.Normal, "Execution query")

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}
