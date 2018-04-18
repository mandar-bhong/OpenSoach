package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/constants"
	"opensoach.com/spl/constants/dbquery"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
)

var SUB_MODULE_NAME = "SPL.User.DB"

func SplMasterUserTableInsert(dbConn string, insrtStruct lmodels.DBSplMasterUserTableRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserTableInsert")

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetSplMasterUserDetailsTableById")

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserDetailsTableInsert")

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserDetailsTableUpdate")

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserTableUpdateState")

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing CheckOldPasswordExists")

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserTableUpdatePassword")

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

func GetSplMasterUserTableTotalFilteredRecords(dbConn string, filterModel *lmodels.DBSearchUserRequestFilterDataModel) (error, *lmodels.DBTotalRecordsModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetSplMasterUserTableTotalFilteredRecords")

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_SPL_MASTER_USER_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter count : "+query)

	selectCtx := dbmgr.SelectContext{}
	data := &lmodels.DBTotalRecordsModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Get()
	if selectErr != nil {
		return selectErr, &lmodels.DBTotalRecordsModel{}
	}
	return nil, data
}

func SplMasterUserTableSelectByFilter(dbConn string, listdatareq lmodels.DataListRequest, filterModel *lmodels.DBSearchUserRequestFilterDataModel, startingRow int) (error, *[]lmodels.DBSearchUserResponseFilterDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserTableSelectByFilter")

	if isParamValid := lhelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		lhelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := lhelper.GetDBTagFromJSONTag(lmodels.DBSearchUserResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_SPL_MASTER_USER_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	query = strings.Replace(query, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter : "+query)

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSearchUserResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(startingRow, limit)
	if selectErr != nil {
		return selectErr, &[]lmodels.DBSearchUserResponseFilterDataModel{}
	}
	return nil, data
}
