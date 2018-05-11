package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/constants"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
)

var SUB_MODULE_NAME = "HKT.API.Servicepoint.DB"

func SpUpdateByFilter(dbConn string, updtStruct *hktmodels.DBSpUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SpUpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_TABLE_SPL_NODE_SP_TBL
	updateErr := updateCtx.UpdateByFilter("SpId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func SpCategoryInsert(dbConn string, insrtStruct *hktmodels.DBSpCategoryInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SpCategoryInsert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_SP_CATEGORY_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func FopSpInsert(dbConn string, insrtStruct *hktmodels.DBFopSpInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing FopSpInsert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_FOP_SP_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func FopSpDelete(dbConn string, deltStruct *lmodels.APIFopSpDeleteRequest) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing FopSpDelete.")

	delDBCtx := dbmgr.UpdateDeleteContext{}
	delDBCtx.DBConnection = dbConn
	delDBCtx.Args = deltStruct
	delDBCtx.QueryType = dbmgr.Query
	delDBCtx.Query = dbquery.QUERY_DELETE_FOP_SP_TABLE_ROW
	deleteErr := delDBCtx.Delete()
	if deleteErr != nil {
		return deleteErr, 0
	}
	return nil, delDBCtx.AffectedRows
}

func SpInsert(dbConn string, insrtStruct *hktmodels.DBSpInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SpInsert.")

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

func GetSpCategoryShortDataList(dbConn string) (error, *[]hktmodels.DBSpCategoryShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetSpCategoryShortDataList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSpCategoryShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_SP_CATEGORY_SHORT_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func DevSpMappingTableInsert(dbConn string, insrtStruct *hktmodels.DBDevSpMappingInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SpDevAssociation.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_NODE_DEV_SP
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func DevSpMappingTableDelete(dbConn string, deltStruct *lmodels.APIDevSpAsscociationRemoveRequest) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SpDevAssociationDelete.")

	delDBCtx := dbmgr.UpdateDeleteContext{}
	delDBCtx.DBConnection = dbConn
	delDBCtx.Args = deltStruct
	delDBCtx.QueryType = dbmgr.Query
	delDBCtx.Query = dbquery.QUERY_DELETE_DEV_SP_MAPPING_TABLE_ROW
	deleteErr := delDBCtx.Delete()
	if deleteErr != nil {
		return deleteErr, 0
	}
	return nil, delDBCtx.AffectedRows
}
