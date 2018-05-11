package dbaccess

import (
	"fmt"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/splserver/constants"
	lmodels "opensoach.com/splserver/models"
)

func UpdateCPMIDToInstDB(dbConn string, insrtStruct *lmodels.APIDBInstanceCpmIdInsertModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing InsertDbInstanceCpmId")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_CPM_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		fmt.Println(insertErr)
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func UpdateDevToInstDB(dbConn string, insrtStruct *lmodels.APIDBInstanceDevInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateDevToInstDB")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_DEV_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		fmt.Println(insertErr)
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}
