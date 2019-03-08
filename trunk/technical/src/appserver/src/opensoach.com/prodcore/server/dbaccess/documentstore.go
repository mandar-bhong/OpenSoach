package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hpft/constants"
	"opensoach.com/prodcore/constants/dbquery"
	pcmodels "opensoach.com/prodcore/models"
)

func InsertDocumentData(dbConn string, insrtStruct *pcmodels.DocumentStoreInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing document insert method.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_HPFT_DOCUMENT_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func GetDocumentByUuid(dbConn string, Uuid string) (error, *[]pcmodels.DocumentStoreRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDocumentByUuid")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]pcmodels.DocumentStoreRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_DOCUMENT_TABLE_SELECT_BY_UUID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(Uuid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func UpdateDocumentData(dbConn string, updtStruct *pcmodels.DocumentStoreUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateDocumentData")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_DOCUMENT_TBL
	updateErr := updateCtx.UpdateByFilter("Uuid", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}
