package dbaccess

import (
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hkthelper "opensoach.com/hkt/api/helper"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
	pcconst "opensoach.com/prodcore/constants"
)

var SUB_MODULE_NAME = "HKT.API.Dashboard.DB"

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

func GetFeedbackSummary(dbConn string, req lmodels.APIDashboardFeedbackRequest) (error, []hktmodels.DBDashBoardFeedbackDataModel) {
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
