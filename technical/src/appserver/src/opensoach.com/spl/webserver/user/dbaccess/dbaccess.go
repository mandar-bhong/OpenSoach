package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/constants"
	"opensoach.com/spl/constants/dbquery"
	lmodels "opensoach.com/spl/models"
)

var SUB_MODULE_NAME = "SPL.User.DB"

func SplMasterUserTableInsert(dbConn string, insrtStruct lmodels.DBSplMasterUserTableRowModel) (error, int64) {
	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = insrtStruct
	insDBCtx.Query = dbquery.QUERY_SPL_MASTER_USER_TABLE_INSERT
	insDBCtx.QueryType = dbmgr.Query
	insDBCtx.TableName = constants.DB_TABLE_USER_TBL
	insErr := insDBCtx.Insert()
	if insErr != nil {
		return insErr, 0
	}
	return nil, insDBCtx.InsertID
}

func GetSplMasterUserDetailsTableById(dbConn string, userid int64) (error, *[]lmodels.DBSplMasterUsrDetailsTableRowModel) {
	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterUsrDetailsTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_MASTER_USR_DETAILS_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(userid)
	if selErr != nil {
		return selErr, &[]lmodels.DBSplMasterUsrDetailsTableRowModel{}
	}
	return nil, data
}

func SplMasterUserDetailsTableInsert(dbConn string, insrtStruct lmodels.DBSplMasterUsrDetailsTableRowModel) (error, int64) {
	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = insrtStruct
	insDBCtx.Query = dbquery.QUERY_SPL_MASTER_USR_DETAILS_TABLE_INSERT
	insDBCtx.QueryType = dbmgr.Query
	insDBCtx.TableName = constants.DB_TABLE_USER_DETAILS_TBL
	insErr := insDBCtx.Insert()
	if insErr != nil {
		return insErr, 0
	}
	return nil, insDBCtx.InsertID
}

func SplMasterUserDetailsTableUpdate(dbConn string, updtStruct lmodels.DBSplMasterUsrDetailsTableRowModel) (error, int64) {
	updtDBCtx := dbmgr.UpdateDeleteContext{}
	updtDBCtx.DBConnection = dbConn
	updtDBCtx.Args = updtStruct
	updtDBCtx.Query = dbquery.QUERY_SPL_MASTER_USR_DETAILS_TABLE_UPDATE
	updtDBCtx.QueryType = dbmgr.Query
	updtDBCtx.TableName = constants.DB_TABLE_USER_DETAILS_TBL
	updateErr := updtDBCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updtDBCtx.AffectedRows
}

func SplMasterUserTableUpdateState(dbConn string, updtStruct lmodels.DBSplMasterUserTableRowModel) (error, int64) {
	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = updtStruct
	updateCtx.Query = dbquery.QUERY_SPL_MASTER_USER_TABLE_UPDATE_STATE
	updateCtx.QueryType = dbmgr.Query
	updateCtx.TableName = constants.DB_TABLE_USER_TBL
	updateErr := updateCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func CheckOldPasswordExists(dbConn string, userid int64, oldPass string) (error, *[]lmodels.DBSplMasterUserTableRowModel) {
	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterUserTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_MASTER_USER_TABLE_SELECT_BY_ID_PASSWORD
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(userid, oldPass)
	if selErr != nil {
		return selErr, &[]lmodels.DBSplMasterUserTableRowModel{}
	}
	return nil, data
}

func SplMasterUserTableUpdatePassword(dbConn string, updtStruct lmodels.DBSplMasterUserTableRowModel) (error, int64) {
	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = updtStruct
	updateCtx.Query = dbquery.QUERY_SPL_MASTER_USER_TABLE_CHANGE_PASSWORD
	updateCtx.QueryType = dbmgr.Query
	updateCtx.TableName = constants.DB_TABLE_USER_TBL
	updateErr := updateCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}
