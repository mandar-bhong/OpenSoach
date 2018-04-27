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

func SplMasterUserTableInsert(dbConn string, insrtStruct lmodels.DBSplMasterUserRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserTableInsert")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
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

func SplMasterUserDetailsTableInsert(dbConn string, insrtStruct lmodels.DBSplMasterUsrDetailsRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserDetailsTableInsert")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_USER_DETAILS_TBL
	insErr := insDBCtx.Insert()
	if insErr != nil {
		return insErr, 0
	}
	return nil, insDBCtx.InsertID
}

func SplMasterUserDetailsTableUpdate(dbConn string, updtStruct lmodels.DBSplMasterUsrDetailsRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserDetailsTableUpdate")

	updtDBCtx := dbmgr.UpdateDeleteContext{}
	updtDBCtx.DBConnection = dbConn
	updtDBCtx.Args = updtStruct
	updtDBCtx.QueryType = dbmgr.AutoQuery
	updtDBCtx.TableName = constants.DB_TABLE_USER_DETAILS_TBL
	updateErr := updtDBCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updtDBCtx.AffectedRows
}

func UpdateUsrState(dbConn string, updtStruct lmodels.DBSplMasterUserRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateUsrState")

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

func CheckOldPasswordExists(dbConn string, userid int64, oldPass string) (error, *[]lmodels.DBSplMasterUserRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing CheckOldPasswordExists")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterUserRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_MASTER_USER_TABLE_SELECT_BY_ID_PASSWORD
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(userid, oldPass)
	if selErr != nil {
		return selErr, &[]lmodels.DBSplMasterUserRowModel{}
	}
	return nil, data
}

func UpdateUsrPassword(dbConn string, updtStruct lmodels.DBSplMasterUserRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateUsrPassword")

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

func GetUserIdByUserName(dbConn string, usrname string) (error, *[]lmodels.DBSplMasterUserTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserIdByUserName")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterUserTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_USERID_BY_USERNAME
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(usrname)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func SplMasterUserCpmTableInsert(dbConn string, insrtStruct lmodels.DBUsrCpmRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterUserTableInsert")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_MASTER_USER_CPM_TBL
	insErr := insDBCtx.Insert()
	if insErr != nil {
		return insErr, 0
	}
	return nil, insDBCtx.InsertID
}

func GetCustUsrFilterList(dbConn string, filterModel *lmodels.DBSearchUserRequestFilterDataModel, listdatareq lmodels.DataListRequest, startingRow int) (error, *lmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUsrFilterList")

	if isParamValid := lhelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		lhelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := lhelper.GetDBTagFromJSONTag(lmodels.DBSearchUserResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_CU_GET_SPL_MASTER_USER_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_CU_SPL_MASTER_USER_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter count : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter : "+listQuery)

	data := &lmodels.ServerListingResultModel{}

	selectCtxCount := dbmgr.SelectContext{}
	dataCount := &lmodels.DBTotalRecordsModel{}
	selectCtxCount.DBConnection = dbConn
	selectCtxCount.Dest = dataCount
	selectCtxCount.Query = countQuery
	selectCtxCount.QueryType = dbmgr.Query
	selectCtxCountErr := selectCtxCount.Get()
	if selectCtxCountErr != nil {
		return selectCtxCountErr, nil
	}

	data.RecordCount = dataCount.TotalRecords

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	resdata := &[]lmodels.DBSearchUserResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = resdata
	selectCtx.Query = listQuery
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(startingRow, limit)
	if selectErr != nil {
		return selectErr, nil
	}

	data.RecordList = resdata

	return nil, data
}

func GetOSUsrFilterList(dbConn string, filterModel *lmodels.DBSearchUserRequestFilterDataModel, listdatareq lmodels.DataListRequest, startingRow int) (error, *lmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUsrFilterList")

	if isParamValid := lhelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		lhelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := lhelper.GetDBTagFromJSONTag(lmodels.DBSearchUserRequestFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_OSU_GET_SPL_MASTER_USER_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_OSU_SPL_MASTER_USER_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter count : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter : "+listQuery)

	data := &lmodels.ServerListingResultModel{}

	selectCtxCount := dbmgr.SelectContext{}
	dataCount := &lmodels.DBTotalRecordsModel{}
	selectCtxCount.DBConnection = dbConn
	selectCtxCount.Dest = dataCount
	selectCtxCount.Query = countQuery
	selectCtxCount.QueryType = dbmgr.Query
	selectCtxCountErr := selectCtxCount.Get()
	if selectCtxCountErr != nil {
		return selectCtxCountErr, nil
	}

	data.RecordCount = dataCount.TotalRecords

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	resdata := &[]lmodels.DBSearchUserResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = resdata
	selectCtx.Query = listQuery
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(startingRow, limit)
	if selectErr != nil {
		return selectErr, nil
	}

	data.RecordList = resdata

	return nil, data
}
