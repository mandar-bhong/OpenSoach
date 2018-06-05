package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
)

var SUB_MODULE_NAME = "HKT.API.Dashboard.DB"

func GetDeviceSummary(dbConn string, cpmid int64) (error, []hktmodels.DBDashBoardDeviceSummaryDataModel) {
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetComplaintList")
	data := []hktmodels.DBDashBoardDeviceSummaryDataModel{}

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = dbquery.QUERY_SPL_NODE_DASHBOARD_DEVICE_SUMMARUY
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select(cpmid)
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}
