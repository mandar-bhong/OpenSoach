package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/splserver/constants/dbquery"
)

const SUB_MODULE_NAME = "SPL.Server.DB"

func GetDBConnectionByID(dbConn string, dbConnID int64) (error, string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDBConnectionByID")

	selDBCtx := dbmgr.SelectContext{}
	data := ""
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_DB_CONN_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = &data
	selErr := selDBCtx.Get(dbConnID)
	if selErr != nil {
		return selErr, ""
	}
	return nil, data
}

func GetDBConnectionByCpmID(dbConn string, cpmID int64) (error, string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDBConnectionByCpmID")

	selDBCtx := dbmgr.SelectContext{}
	data := ""
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_DB_CONN_BY_CPM_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = &data
	selErr := selDBCtx.Get(cpmID)
	if selErr != nil {
		return selErr, ""
	}
	return nil, data
}
