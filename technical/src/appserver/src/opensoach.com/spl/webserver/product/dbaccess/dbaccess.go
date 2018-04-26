package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/constants/dbquery"
	lmodels "opensoach.com/spl/models"
)

var SUB_MODULE_NAME = "SPL.Product.DB"

func DBSplMasterProductTableRowModelSelectAll(dbConn string) (error, *[]lmodels.DBSplMasterProductTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetProdList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterProductTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_MASTER_PRODUCT_TBL_SELECT_ALL
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetDbinstanceList(dbConn string) (error, *[]lmodels.DBSplMasterDatabaseInstanceTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDbinstanceList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterDatabaseInstanceTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_MASTER_DATABASE_INSTANCE_SELECT_ALL
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
