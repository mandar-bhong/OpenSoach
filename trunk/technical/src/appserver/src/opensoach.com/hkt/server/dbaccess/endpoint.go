package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"
	hktmodels "opensoach.com/hkt/models"
	prodmodels "opensoach.com/hkt/models"
	"opensoach.com/hkt/server/constants"
	"opensoach.com/hkt/server/constants/dbquery"
)

func EPGetInstanceDB(dbConn string, cpmid, deviceid int64) (error, string) {
	selDBCtx := dbmgr.SelectContext{}
	data := ""
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_DB_INST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = &data
	selErr := selDBCtx.Get(cpmid)
	if selErr != nil {
		return selErr, ""
	}

	return nil, data
}

func EPGetDeviceServicePoints(dbConn string, cpmid int64, deviceid int64) (error, *[]prodmodels.DBEPSPDataModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]prodmodels.DBEPSPDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_DEVICE_SP
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid, deviceid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPGetSPAuthCodes(dbConn string, cpmid int64, spid int64) (error, *[]string) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]string{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_SP_OPERATOR
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid, spid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPGetSPServConf(dbConn string, cpmid int64, spid int64) (error, *[]prodmodels.DBEPSPServConfDataModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]prodmodels.DBEPSPServConfDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_SP_SERV_CONF
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid, spid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPInsertServiceInstanceData(dbConn string,
	dbSerInstDataRowModel hktmodels.DBServiceInstanceDataRowModel) error {

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = dbSerInstDataRowModel
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_NODE_SERVICE_INST_TXN_TBL
	insertErr := insDBCtx.Insert()

	return insertErr
}
