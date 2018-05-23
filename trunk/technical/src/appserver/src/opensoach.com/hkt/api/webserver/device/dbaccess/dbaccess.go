package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
)

var SUB_MODULE_NAME = "HKT.API.Device.DB"

func GetDeviceShortDataList(dbConn string) (error, *[]hktmodels.DBDeviceShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceShortDataList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBDeviceShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_NODE_DEVICE_TABLE_SELECT_SHORT_DATA_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
