package dbaccess

import (
	"time"

	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/prodcore/constants/dbquery"
	"opensoach.com/prodcore/models"
)

func GetSyncConfig(dbConn string, storename string, devicetype int) (error, *models.SyncConfigModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &models.SyncConfigModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_SYNC_CONFIG_ON
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(storename, devicetype)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetTableData(dbConn string, query string, params interface{}) (error, []map[string]interface{}) {

	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = query
	selDBCtx.QueryType = dbmgr.Query
	selErr, data := selDBCtx.SelectToMap(params)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetTableDataCount(dbConn string, query string, params interface{}) (error, *models.SyncConfigTblInfoModel) {

	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	data := &models.SyncConfigTblInfoModel{}
	selDBCtx.Dest = data
	selDBCtx.Query = query
	selDBCtx.QueryType = dbmgr.Query
	selErr, countresult := selDBCtx.SelectToMap(params)
	if selErr != nil {
		return selErr, nil
	}

	data.Count = int((countresult[0]["count"]).(int64))

	if countresult[0]["max_updated_on"] != nil {
		maxUpdatedOn := countresult[0]["max_updated_on"].(time.Time)
		data.MaxUpdatedOn = &maxUpdatedOn
	} else {
		data.MaxUpdatedOn = nil
	}

	return nil, data
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
