package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hkthelper "opensoach.com/hkt/api/helper"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/constants"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
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

func GetServicePointList(dbConn string, filterModel *hktmodels.DBSearchServicePointRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetServicePointList")

	if isParamValid := hkthelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		hkthelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := hkthelper.GetDBTagFromJSONTag(hktmodels.DBSearchServicePointResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := hkthelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_GET_NODE_SP_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_NODE_SP_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Point Filter Record list filter count query : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Point Filter Record list filter query : "+listQuery)

	data := &gmodels.ServerListingResultModel{}

	selectCtxCount := dbmgr.SelectContext{}
	dataCount := &hktmodels.DBTotalRecordsModel{}
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
	resdata := &[]hktmodels.DBSearchServicePointResponseFilterDataModel{}
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

func FopSpSelectByID(dbConn string, fopId int64) (error, *[]hktmodels.DBFopSpAssociationDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing FopSpSelectByID")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBFopSpAssociationDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Query = dbquery.QUERY_GET_FOP_SP_ASSOCIATIONS
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(fopId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetServicePointShortDataList(dbConn string) (error, *[]hktmodels.DBServicePointShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetSpCategoryShortDataList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBServicePointShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_SERVICEPOINT_SHORT_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
