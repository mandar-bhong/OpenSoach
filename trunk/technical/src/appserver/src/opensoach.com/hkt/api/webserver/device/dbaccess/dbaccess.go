package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hkthelper "opensoach.com/hkt/api/helper"
	"opensoach.com/hkt/constants"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Device.DB"

func GetDeviceShortDataList(dbConn string, cpmId int64) (error, *[]hktmodels.DBDeviceShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceShortDataList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBDeviceShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_NODE_DEVICE_TABLE_SELECT_SHORT_DATA_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetDeviceList(dbConn string, filterModel *hktmodels.DBSearchDeviceRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceList")

	if isParamValid := hkthelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		hkthelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := hkthelper.GetDBTagFromJSONTag(hktmodels.DBSearchDeviceResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := hkthelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_GET_DEVICE_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_DEVICE_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device Filter Record list filter count query : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device Filter Record list filter query : "+listQuery)

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
	resdata := &[]hktmodels.DBSearchDeviceResponseFilterDataModel{}
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

func GetDeviceWithNoSpAssociationShortDataList(dbConn string, cpmid int64) (error, *[]hktmodels.DBDeviceShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceWithNoSpAssociationShortDataList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBDeviceShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_DEVICE_WITH_NO_SP_ASSOCIATION_SHORT_DATA_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetDeviceById(dbConn string, devid int64) (error, *[]hktmodels.DBSplNodeDevTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplNodeDevTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_NODE_DEVICE_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(devid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func UpdateByFilter(dbConn string, updtStruct *hktmodels.DBDeviceUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Device UpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_DEV_TBL
	updateErr := updateCtx.UpdateByFilter("DevId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}
