package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	gmodels "opensoach.com/models"
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

func GetCustUsrFilterList(dbConn string, filterModel *lmodels.DBSearchUserRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetCustUsrFilterList")

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

	data := &gmodels.ServerListingResultModel{}

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

func GetOSUsrFilterList(dbConn string, filterModel *lmodels.DBSearchUserRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetOSUsrFilterList")

	if isParamValid := lhelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		lhelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := lhelper.GetDBTagFromJSONTag(lmodels.DBSearchUserResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_OSU_GET_SPL_MASTER_USER_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_OSU_SPL_MASTER_USER_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter count : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter : "+listQuery)

	data := &gmodels.ServerListingResultModel{}

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

func GetUserById(dbConn string, userId int64) (error, *[]lmodels.DBUserInfoDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBUserInfoDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_USER_TABLE_INFO_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(userId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetUserDetailsById(dbConn string, userId int64) (error, *[]lmodels.DBSplMasterUsrDetailsTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserDetailsById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterUsrDetailsTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_SPL_MASTER_USER_DETAILS_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(userId)
	if selErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while get customer id .", selErr)
		return selErr, nil
	}
	return nil, data
}

func GetUroleListOSU(dbConn string) (error, *[]lmodels.DBUroleShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUroleListOSU")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBUroleShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_UROLE_LIST_OSU
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetUroleList(dbConn string, productCode string) (error, *[]lmodels.DBUroleShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUroleListCU")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBUroleShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_UROLE_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(productCode)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetProdAssociationByUsrId(dbConn string, usrId int64) (error, *[]lmodels.DBUserProdAssociationDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetProdAssociationByUsrId")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBUserProdAssociationDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_PRODUCT_ASSOCIATION_BY_USER_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(usrId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func UcpmStateUpdate(dbConn string, updtStruct *lmodels.DBUsrCpmStateUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UcpmStateUpdate")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_TABLE_MASTER_USER_CPM_TBL
	updateErr := updateCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func UserUpdate(dbConn string, updtStruct *lmodels.DBUserUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UserUpdate")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_TABLE_USER_TBL
	updateErr := updateCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}
