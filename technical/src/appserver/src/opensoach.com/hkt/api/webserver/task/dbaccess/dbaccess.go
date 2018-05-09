package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hkt/constants"
	hktmodels "opensoach.com/hkt/models"
)

var SUB_MODULE_NAME = "HKT.API.Task.DB"

func Insert(dbConn string, insrtStruct *hktmodels.DBTaskLibInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Task insert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_HKT_TASK_LIB
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func UpdateByFilter(dbConn string, updtStruct *hktmodels.DBTaskLibUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Task UpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_TABLE_HKT_TASK_LIB
	updateErr := updateCtx.UpdateByFilter("TaskLibId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func TaskSelectByID(dbConn string, taskId int64) (error, *[]hktmodels.DBSplHktTaskLibTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing TaskSelectByID")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHktTaskLibTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_TABLE_HKT_TASK_LIB
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(taskId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func TaskSelectAll(dbConn string) (error, *[]hktmodels.DBSplHktTaskLibTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing TaskSelectAll")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHktTaskLibTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_TABLE_HKT_TASK_LIB
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectAll()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data

}
