package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hkt/constants"
	hktmodels "opensoach.com/hkt/models"
)

var SUB_MODULE_NAME = "HKT.API.Master.DB"

func GetHktMasterSpcTaskLibTableById(dbConn string, id int64) (error, *[]hktmodels.DBSplHktMasterSpcTaskLibTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Field GetHktMasterSpcTaskLibTableById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHktMasterSpcTaskLibTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_TABLE_HKT_MASTER_SPC_TASK_LIB_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(id)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetHktMasterTaskLibTableById(dbConn string, id int64) (error, *[]hktmodels.DBSplHktMasterTaskLibTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Field GetHktMasterTaskLibTableById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHktMasterTaskLibTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_TABLE_HKT_MASTER_TASK_LIB_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(id)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetSplProdMasterServConfTypeTableById(dbConn string, id int64) (error, *[]hktmodels.DBSplProdMasterServConfTypeTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Field GetSplProdMasterServConfTypeTableById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplProdMasterServConfTypeTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_PROD_MASTER_SERV_CONF_TYPE_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(id)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetSplProdMasterSpCategoryTableById(dbConn string, id int64) (error, *[]hktmodels.DBSplProdMasterSpCategoryTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Field GetSplProdMasterSpCategoryTableById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplProdMasterSpCategoryTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_PROD_MASTER_SP_CATEGORY_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(id)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
