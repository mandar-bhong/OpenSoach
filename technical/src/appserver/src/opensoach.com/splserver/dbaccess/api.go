package dbaccess

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/splserver/constants"
	lmodels "opensoach.com/splserver/models"
)

func GetDBTransaction(dbconn string) (error, *sqlx.Tx) {
	ctx := dbmgr.InsertTxContext{}
	return ctx.GetTransaction(dbconn)
}

func UpdateCPMIDToInstDB(tx *sqlx.Tx, insrtStruct *lmodels.APITaskDBInstanceCpmIdInsertModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing InsertDbInstanceCpmId")

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
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

func UpdateDevToInstDB(tx *sqlx.Tx, insrtStruct *lmodels.APITaskDBInstanceDevInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateDevToInstDB")

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_DEV_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func UpdateSpCategoryToInstanceDB(tx *sqlx.Tx, insrtStruct *lmodels.APITaskDBInstanceSpCategoryInsertModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateSpCategoryToInstanceDB")

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_SP_CATEGORY_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func UpdateServicePointsToInstDB(dbConn string, insrtStruct *lmodels.APITaskDBNodeSpInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateServicePointsToInstDB")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_SP_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func UpdateDevStatusToInstDB(tx *sqlx.Tx, insrtStruct *lmodels.APITaskDBInstanceDevStatusInsertModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateDevStatusToInstDB")

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_DEV_STATUS_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}
