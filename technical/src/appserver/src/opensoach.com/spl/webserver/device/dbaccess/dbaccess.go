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

var SUB_MODULE_NAME = "SPL.Device.DB"

func SplMasterDeviceTableInsert(dbConn string, insrtStruct *lmodels.DBSplMasterDeviceRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterDeviceTableInsert")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_MASTER_DEVICE_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func UpdateDeviceState(dbConn string, updtStruct *lmodels.DevStateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateDeviceState")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_TABLE_MASTER_DEVICE_TBL
	updateErr := updateCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func GetDeviceId(dbConn string, cpmid int64) (error, *lmodels.DBSplMasterCpmDevMappingTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceId")

	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBSplMasterCpmDevMappingTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_DEV_ID_BY_CPM_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(cpmid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetSplMasterDeviceDetailsTableById(dbConn string, devid int64) (error, *[]lmodels.DBSplMasterDevDetailsTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetSplMasterDeviceDetailsTableById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterDevDetailsTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_MASTER_DEVICE_DETAILS_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(devid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func SplMasterDeviceDetailsTableInsert(dbConn string, insrtStruct *lmodels.DBSplMasterDevDetailsRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterDeviceDetailsTableInsert")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_MASTER_DEVICE_DETAILS_TBL
	insErr := insDBCtx.Insert()
	if insErr != nil {
		return insErr, 0
	}
	return nil, insDBCtx.InsertID
}

func SplMasterDeviceDetailsTableUpdate(dbConn string, updtStruct *lmodels.DBSplMasterDevDetailsRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterDeviceDetailsTableUpdate")

	updtDBCtx := dbmgr.UpdateDeleteContext{}
	updtDBCtx.DBConnection = dbConn
	updtDBCtx.Args = *updtStruct
	updtDBCtx.QueryType = dbmgr.AutoQuery
	updtDBCtx.TableName = constants.DB_TABLE_MASTER_DEVICE_DETAILS_TBL
	updateErr := updtDBCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updtDBCtx.AffectedRows
}

func GetDeviceFilterRecordsCount(dbConn string, filterModel *lmodels.DBSearchDeviceRequestFilterDataModel) (error, *lmodels.DBTotalRecordsModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceFilterRecordsCount")

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_SPL_MASTER_DEVICE_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device Filter Record list filter count query : "+query)

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

func GetDeviceListData(dbConn string, listdatareq lmodels.DataListRequest, filterModel *lmodels.DBSearchDeviceRequestFilterDataModel, startingRow int) (error, *[]lmodels.DBSearchDeviceResponseFilterDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceListData")

	if isParamValid := lhelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		lhelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := lhelper.GetDBTagFromJSONTag(lmodels.DBSearchDeviceRequestFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_SPL_MASTER_DEVICE_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	query = strings.Replace(query, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device Filter Record list filter query : "+query)

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSearchDeviceResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(startingRow, limit)
	if selectErr != nil {
		return selectErr, &[]lmodels.DBSearchDeviceResponseFilterDataModel{}
	}
	return nil, data
}

func SetDeviceCustId(dbConn string, updtStruct *lmodels.DevCustRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SetDeviceCustId")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_TABLE_MASTER_DEVICE_TBL
	updateErr := updateCtx.Update()
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func CpmDevTableInsert(dbConn string, insrtStruct *lmodels.DBSplCpmDevRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing CpmDevTableInsert")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_MASTER_CPM_DEV_MAPPING_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}
