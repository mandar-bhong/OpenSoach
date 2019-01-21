package dbaccess

import (
	"time"

	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/prodcore/constants/dbquery"
	"opensoach.com/prodcore/models"
)

func GetSyncConfig(dbConn string, storename string) (error, *models.SyncConfigModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &models.SyncConfigModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_SYNC_CONFIG_ON
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(storename)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetTableData(dbConn string, query string, updateon time.Time) (error, []map[string]interface{}) {

	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = query
	selDBCtx.QueryType = dbmgr.Query
	selErr, data := selDBCtx.SelectToMap(updateon)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetTableDataCount(dbConn string, query string, updateon time.Time) (error, int) {

	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	var count int
	selDBCtx.Dest = &count
	selDBCtx.Query = query
	selDBCtx.QueryType = dbmgr.Query
	selErr := selDBCtx.Get(updateon)
	if selErr != nil {
		return selErr, 0
	}
	return nil, count
}

func GetTableDataByUuid(dbConn string, query string, uuid string) (error, int) {

	selDBCtx := dbmgr.SelectContext{}
	var count int
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = query
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = &count
	selErr := selDBCtx.Get(uuid)
	if selErr != nil {
		return selErr, 0
	}
	return nil, count
}

func InsertTableData(dbConn string, query string, insrtmodel interface{}) error {

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = insrtmodel
	insDBCtx.Query = query
	insDBCtx.QueryType = dbmgr.Query
	insertErr := insDBCtx.Insert()

	return insertErr
}

func UpdateTableData(dbConn string, query string, updtStruct interface{}) (error, int64) {

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = updtStruct
	updateCtx.Query = query
	updateCtx.QueryType = dbmgr.Query
	updateErr := updateCtx.UpdateByFilter("Uuid")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}
