package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hkt/constants"
	hktmodels "opensoach.com/hkt/models"
)

var SUB_MODULE_NAME = "HKT.Task.DB"

func Insert(dbConn string, insrtStruct *hktmodels.DBTaskLibRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Task insert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_TASK_LIB
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func Update(dbConn string) {

}

func SelectByID(dbConn string) {

}

func SelectAll(dbConn string) {

}
