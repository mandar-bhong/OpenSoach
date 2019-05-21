package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hpfthelper "opensoach.com/hpft/api/helper"
	"opensoach.com/hpft/constants"
	"opensoach.com/hpft/constants/dbquery"
	hpftmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.Endpoint.DB"

func GetPatientList(dbConn string, usrid int64, filterModel *hpftmodels.DBDeviceSearchPatientRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientList")

	if isParamValid := hpfthelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		hpfthelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	filterConfigModel := gmodels.FilterConfigModel{}
	filterConfigModel.OrAndOperator = " OR "

	dbMatchedTag := hpfthelper.GetDBTagFromJSONTag(hpftmodels.DBDeviceSearchPatientResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := hpfthelper.GetFilterConditionConfigFormModel(*filterModel, filterConfigModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_GET_DEEVICE_SPL_PATIENT_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_DEVICE_SPL_PATIENT_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient Filter Record list filter count query : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient Filter Record list filter query : "+listQuery)

	data := &gmodels.ServerListingResultModel{}

	selectCtxCount := dbmgr.SelectContext{}
	dataCount := &hpftmodels.DBTotalRecordsModel{}
	selectCtxCount.DBConnection = dbConn
	selectCtxCount.Dest = dataCount
	selectCtxCount.Query = countQuery
	selectCtxCount.QueryType = dbmgr.Query
	selectCtxCountErr := selectCtxCount.Get(usrid, usrid)
	if selectCtxCountErr != nil {
		return selectCtxCountErr, nil
	}

	data.RecordCount = dataCount.TotalRecords

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	resdata := &[]hpftmodels.DBDeviceSearchPatientResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = resdata
	selectCtx.Query = listQuery
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(usrid, usrid, startingRow, limit)
	if selectErr != nil {
		return selectErr, nil
	}

	data.RecordList = resdata

	return nil, data
}

func PatientUserAssociation(dbConn string, insrtStruct *hpftmodels.DBPatientMonitorMappingInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing PatientUserAssociation.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_HPFT_USER_PATIENT_MONITOR_MAPPING
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func PatientUserDeAssociation(dbConn string, deltStruct *hpftmodels.DBPatientMonitorMappingDeleteRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing PatientUserDeAssociation.")

	delDBCtx := dbmgr.UpdateDeleteContext{}
	delDBCtx.DBConnection = dbConn
	delDBCtx.Args = deltStruct
	delDBCtx.QueryType = dbmgr.Query
	delDBCtx.TableName = constants.DB_SPL_HPFT_USER_PATIENT_MONITOR_MAPPING
	delDBCtx.Query = dbquery.QUERY_DELETE_USER_PATIENT_ASSOCIATION

	if deltStruct.UsrId == nil {
		delDBCtx.Query = dbquery.QUERY_DELETE_USER_PATIENT_ASSOCIATION_BY_PATIENT_ID
	} else {
		if deltStruct.SpId == nil {
			delDBCtx.Query = dbquery.QUERY_DELETE_USER_PATIENT_ASSOCIATION_BY_USER_ID
		}

		if deltStruct.SpId != nil && deltStruct.PatientId == nil {
			delDBCtx.Query = dbquery.QUERY_DELETE_USER_PATIENT_ASSOCIATION_BY_SP
		}
	}

	deleteErr := delDBCtx.Delete()
	if deleteErr != nil {
		return deleteErr, 0
	}
	return nil, delDBCtx.AffectedRows
}

func GetUserPatientassociationByUsrId(dbConn string, usrid int64) (error, *[]hpftmodels.DBSplHpftUserPatientMonitorMappingRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserPatientassociationByUsrId")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hpftmodels.DBSplHpftUserPatientMonitorMappingRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_USER_PATIENT_ASSOCIATION_BY_USER_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(usrid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetUserPatientassociationByUsrIdSpId(dbConn string, usrid, spid int64) (error, *[]hpftmodels.DBSplHpftUserPatientMonitorMappingRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserPatientassociationByUsrIdSpId")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hpftmodels.DBSplHpftUserPatientMonitorMappingRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_USER_PATIENT_ASSOCIATION_BY_USER_ID_SP_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(usrid, spid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
