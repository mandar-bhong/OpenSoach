package dbaccess

import (
	"fmt"
	"strings"

	"opensoach.com/core/logger"

	"errors"

	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/constants"
	"opensoach.com/spl/constants/dbquery"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
)

const SUB_MODULE_NAME = "SPL.Customer.DB"

func GetCustomerById(dbConn string, customerId int64) (error, *[]lmodels.DBSplMasterCustomerTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetCustomerById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterCustomerTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_CUSTOMER_TABLE_INFO_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(customerId)
	if selErr != nil {
		return selErr, &[]lmodels.DBSplMasterCustomerTableRowModel{}
	}
	return nil, data
}

func GetCustomerDetailsById(dbConn string, customerId int64) (error, *[]lmodels.DBSplMasterCustDetailsTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetCorpDetailsById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterCustDetailsTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_MASTER_CUST_DETAILS_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(customerId)
	if selErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while get customer id .", selErr)
		return selErr, nil
	}
	return nil, data
}

func GetCorpDetailsById(dbConn string, customerId int64) (error, *lmodels.DBSplMasterCorpTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetCorpDetailsById")

	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBSplMasterCorpTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_CORP_TABLE_INFO_BY_CUSTOMER_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(customerId)
	if selErr != nil {
		return selErr, &lmodels.DBSplMasterCorpTableRowModel{}
	}
	return nil, data
}

func GetSplMasterCustomerTableTotalFilteredRecords(dbConn string, filterModel *lmodels.DBSearchCustomerRequestFilterDataModel) (error, *lmodels.DBTotalRecordsModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetSplMasterCustomerTableTotalFilteredRecords")

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_SPL_MASTER_CUSTOMER_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Customer Filter Record list filter count : "+query)

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

func SplMasterCustomerTableSelectByFilter(dbConn string, listdatareq lmodels.DataListRequest, filterModel *lmodels.DBSearchCustomerRequestFilterDataModel, startingRow int) (error, *[]lmodels.DBSearchCustomerResponseFilterDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterCustomerTableSelectByFilter")

	if isParamValid := lhelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		lhelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := lhelper.GetDBTagFromJSONTag(lmodels.DBSearchCustomerRequestFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_SPL_MASTER_CUSTOMER_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	query = strings.Replace(query, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Customer Filter Record list filter : "+query)

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSearchCustomerResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(startingRow, limit)
	if selectErr != nil {
		return selectErr, &[]lmodels.DBSearchCustomerResponseFilterDataModel{}
	}
	return nil, data
}

func GetCustList(dbConn string, filterModel *lmodels.DBSearchCustomerRequestFilterDataModel, listdatareq lmodels.DataListRequest, startingRow int) (error, *lmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetCustList")

	if isParamValid := lhelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		lhelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := lhelper.GetDBTagFromJSONTag(lmodels.DBSearchCustomerResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_GET_SPL_MASTER_CUSTOMER_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_SPL_MASTER_CUSTOMER_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Customer Filter Record list filter count query : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Customer Filter Record list filter query : "+listQuery)

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
	resdata := &[]lmodels.DBSearchCustomerResponseFilterDataModel{}
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

func AddCustomer(dbconn string, req lmodels.DBSplMasterCustomerTableRowModel) (error, int64) {

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbconn
	insDBCtx.Query = dbquery.QUERY_SPL_MASTER_CUSTOMER_TABLE_INSERT
	insDBCtx.QueryType = dbmgr.Query
	insDBCtx.Args = req

	intErr := insDBCtx.Insert()

	if intErr != nil {
		return intErr, 0
	}

	return nil, insDBCtx.InsertID
}

func CustomerDetailsTableInsert(dbconn string, req lmodels.DBSplMasterCustDetailsTableRowModel) (error, int64) {

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbconn
	insDBCtx.Query = dbquery.QUERY_SPL_MASTER_CUST_DETAILS_TABLE_INSERT
	insDBCtx.QueryType = dbmgr.Query
	insDBCtx.Args = req

	insErr := insDBCtx.Insert()

	if insErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while insert customer details .", insErr)

		return insErr, 0
	}

	return nil, insDBCtx.InsertID
}

func CustomerDetailsTableUpdate(dbconn string, req lmodels.DBSplMasterCustDetailsTableRowModel) (error, int64) {

	updDBCtx := dbmgr.UpdateDeleteContext{}
	updDBCtx.DBConnection = dbconn
	updDBCtx.Query = dbquery.QUERY_SPL_MASTER_CUST_DETAILS_TABLE_UPDATE
	updDBCtx.QueryType = dbmgr.Query
	updDBCtx.Args = req

	updErr := updDBCtx.Update()

	if updErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while update customer details .", updErr)
		return updErr, 0
	}

	return nil, updDBCtx.AffectedRows

}

func CpmTableInsert(dbConn string, insrtStruct *lmodels.DBCustProdMappingInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing CpmTableInsert")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_MASTER_CUST_PROD_MAPPING_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}
