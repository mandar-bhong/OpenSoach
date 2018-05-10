package dbaccess

import (
	"opensoach.com/splserver/constants/dbquery"
)

const SUB_MODULE_NAME = "SPL.Server.DB"

func GetDBConnectionByID(dbconn string, dbConnID int) (error, string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetCustomerById")

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
