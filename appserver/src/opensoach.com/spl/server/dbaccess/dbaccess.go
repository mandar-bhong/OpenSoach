package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hktmodels "opensoach.com/hkt/models"
	"opensoach.com/spl/server/constants/dbquery"
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

func GetDBHktMasterSpCategory(dbConn string) (error, *[]hktmodels.DBSplProdMasterSpCategoryTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDBHktMasterSpCategory")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplProdMasterSpCategoryTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_ALL_PROD_MASTER_SP_CATEGORY_TBL
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetDBHktMasterTaskLib(dbConn string) (error, *[]hktmodels.DBSplHktMasterTaskLibTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDBHktMasterTaskLib")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHktMasterTaskLibTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_ALL_HKT_MASTER_TASK_LIB_TBL
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetUserName(dbConn string, userid int64) (error, string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserName")

	selDBCtx := dbmgr.SelectContext{}
	data := ""
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_USER_USERNAME_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = &data
	selErr := selDBCtx.Get(userid)
	if selErr != nil {
		return selErr, ""
	}
	return nil, data
}

func GetProdCodeByProdID(dbConn string, prodID int64) (error, string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetProdCodeByProdID")

	selDBCtx := dbmgr.SelectContext{}
	data := ""
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_PROD_CODE_BY_PROD_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = &data
	selErr := selDBCtx.Get(prodID)
	if selErr != nil {
		return selErr, ""
	}
	return nil, data
}
